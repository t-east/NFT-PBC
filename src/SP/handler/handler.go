
package handler

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	"fmt"
	"bytes"
	// "encoding/binary"
	"pairing_test/src/tool"
	// "reflect"
	"os"
	"encoding/json"
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
			ArtId: requestBody.ArtId}
		artLog := structure.ArtLog{
			HashedData:  requestBody.HashedData,
			Owner:  requestBody.Owner,}
		
		artId := requestBody.ArtId
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
        var response structure.HashedResponseToUser
		response.HashedFile = GetHashDataFromBN(artId)
		response.ArtId = artId
		log.Print(response)
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, response)
    }
}

// ファイルログをブロックチェーンに保存
func SaveOutsourceData(artLog structure.ArtLog, artId string) error {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal("get-go-env-error")
		}
	privKey := os.Getenv("SP_PRIVATE_KEY")
	conn, client := ethhandler.ConnectNetWork()
	auth := ethhandler.AuthUser(client, privKey)
	userAddress := ethhandler.StringToAddress(artLog.Owner)
	reply, err := conn.RegisterArt(auth, artLog.HashedData, artId , userAddress)
	if err != nil{
		log.Fatal(err)
	}
	log.Print(reply)
	return err
}

// Artログをブロックチェーンから抜き出し
func GetHashDataFromBN(artId string) [][]byte {
	conn, _ := ethhandler.ConnectNetWork()	
	reply := ethhandler.GetHashData(conn, artId)
	return reply
}

// 監査証明作成
func AuditProofGen(para *structure.Params, storage *structure.Storage) gin.HandlerFunc {
    return func(c *gin.Context) {
		challens := []structure.Chal{}
		c.Bind(&challens)
		var myu *pbc.Element
		var gamma *pbc.Element
		proofs := &structure.Proofs{}
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for i := 0; i < len(storage.Datas); i++ {
			var chal structure.Chal
			for j := 0; j < len(challens); j++ {
				if storage.Datas[i].ArtId == challens[j].ArtId {
					chal = challens[j]
					break
				}
			}
			splitedFile, _ := tool.SplitSlice(storage.Datas[i].File, len(storage.Datas[i].MetaData))
			aTable, vTable := tool.HashChallen(storage.Datas[i].SplitCount, chal.C, chal.K1, chal.K2, pairing)
			for cIndex := 0; cIndex < chal.C; cIndex++ {
				meta := pairing.NewG1().SetBytes(storage.Datas[i].MetaData[aTable[cIndex]])
				m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
				if cIndex == 0 {
					myu = pairing.NewZr().MulBig(vTable[cIndex], m.X())
					gamma = pairing.NewG1().PowZn(meta, vTable[cIndex])
				} else {
					myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[cIndex], m.X()))
					gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[cIndex]))
				}
			}
			proof := structure.Proof{
				Myu:  myu.Bytes(),
				Gamma:  gamma.Bytes(),
				ArtId: storage.Datas[i].ArtId}
			proofs.AddProofs(proof)
		}
		reqJson := new(bytes.Buffer)
		log.Print(proofs)
		json.NewEncoder(reqJson).Encode(proofs)
		res, err := http.Post("http://tpa:4002/audit", "application/json", reqJson)
		if err != nil {
			log.Fatal("[!] " + err.Error())
		} else {
			log.Print("[*] " + res.Status)
		}
		c.JSON(http.StatusOK, proofs)
	}
}
