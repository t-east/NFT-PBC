
package handler

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	"fmt"
	// "encoding/binary"
	"pairing_test/src/tool"
	"reflect"
	"os"
	"github.com/joho/godotenv"
	"pairing_test/src/SP/structure"
	"pairing_test/src/ethereum/ethhandler"

	"github.com/Nik-U/pbc"
)

func GetPara(para *structure.Params) {
	conn, _ := ethhandler.ConnectNetWork()
	reply := ethhandler.GetPara(conn)
	para.G = reply.G
	para.U = reply.U
	para.Pairing = reply.Pairing
}

//　ファイルをストレージに保存
func ArtPost(storage *structure.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
        requestBody := structure.UploadFile{}
        c.BindJSON(&requestBody)
		log.Print(requestBody.HashedData)
        uploadItem := structure.UploadData{
			File: requestBody.File,
			MetaData: requestBody.MetaData,
			FileName: requestBody.FileName,
			Owner: requestBody.Owner,
			SplitCount: requestBody.SplitCount,
			ArtId: []byte(requestBody.ArtId)}
		artLog := structure.ArtLog{
			HashedData:  requestBody.HashedData,
			Owner:  requestBody.Owner,}
		
		artId := []byte(requestBody.ArtId)
		err := SaveOutsourceData(artLog, artId)
		if err != nil{
			log.Fatal(err)
		}
        storage.AddStorage(uploadItem)
		log.Print(storage)
        c.Status(http.StatusNoContent)
		log.Print(artId)
		c.JSON(http.StatusOK, artLog)
    }
}

//　ファイルをストレージに保存
func ArtGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		artId := c.Param("id")
		artIdByte := []byte(artId)
        var response structure.HashedResponseToUser
		response.HashedFile = GetHashDataFromBN(artIdByte)
		response.ArtId = artId
		log.Print(response)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, response)
    }
}

// ファイルログをブロックチェーンに保存
func SaveOutsourceData(artLog structure.ArtLog, artId []byte) error {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal("get-go-env-error")
		}
	privKey := os.Getenv("SP_PRIVATE_KEY")
	err = ethhandler.RegisterArt(privKey, artLog.Owner, artLog.HashedData, artId)
	return err
}

// Artログをブロックチェーンから抜き出し
func GetHashDataFromBN(artId []byte) [][]byte {
	reply := ethhandler.GetHashData(artId)
	return reply
}

// 監査証明作成
func AuditProofGen(para *structure.Params, storage *structure.Storage) gin.HandlerFunc {
    return func(c *gin.Context) {
		logs := ethhandler.GetLog()
		var myu *pbc.Element
		var gamma *pbc.Element
		proofs := &structure.Proofs{}
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for i := 0; i < len(logs); i++ {
			var storageItem structure.UploadData
			for j := 0; j < len(storage.Datas); j++ {
				if reflect.DeepEqual(storage.Datas[j].ArtId, logs[i].Log.ArtId) {
					storageItem = storage.Datas[j]
					break
				}
			}
			splitedFile, _ := tool.SplitSlice(storageItem.File, len(storageItem.MetaData))
			aTable, vTable := tool.HashChallen(storageItem.SplitCount, int(logs[i].Log.Chal), logs[i].Log.K1, logs[i].Log.K2, pairing)
			for cIndex := 0; cIndex < int(logs[i].Log.Chal); cIndex++ {
				meta := pairing.NewG1().SetBytes(storageItem.MetaData[aTable[cIndex]])
				m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
				if cIndex == 0 {
					myu = pairing.NewZr().MulBig(vTable[cIndex], m.X())
					gamma = pairing.NewG1().PowZn(meta, vTable[cIndex])
				} else {
					myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[cIndex], m.X()))
					gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[cIndex]))
				}
			}
			err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
			if err != nil {
				log.Fatal("get-go-env-error")
			}
			privKey := os.Getenv("SP_PRIVATE_KEY")
			err = ethhandler.SetProof(privKey, myu.Bytes(), gamma.Bytes(), logs[i].Log.ArtId, logs[i].LogId)
			if err != nil {
				log.Fatal("couldnt set proof")
			}
		}
		c.JSON(http.StatusOK, proofs)
	}
}
