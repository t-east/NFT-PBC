
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
	reply := ethhandler.GetParam(conn)
	para.G = reply.G
	para.U = reply.U
	para.Pairing = reply.Pairing
}

//　ファイルをストレージに保存
func ArtPost(storage *structure.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
        requestBody := structure.UploadFile{}
        c.BindJSON(&requestBody)
        uploadItem := structure.UploadData{
			File: requestBody.File,
			MetaData: requestBody.MetaData,
			FileName: requestBody.FileName,
			Owner: requestBody.Owner,
			SplitCount: requestBody.SplitCount,
			ArtId: []byte(requestBody.ArtId),
			HashedData:  requestBody.HashedData,
		}
		artLog := structure.ArtLog{
			HashedData:  requestBody.HashedData,
			Owner:  requestBody.Owner,}
		
		artId := []byte(requestBody.ArtId)
		err := SaveOutsourceData(artLog, artId)
		if err != nil{
			log.Fatal(err)
		}
        storage.AddStorage(uploadItem)
        c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, artLog)
    }
}

// ファイルログをブロックチェーンに保存
func SaveOutsourceData(artLog structure.ArtLog, artId []byte) error {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal("get-go-env-error")
		}
	privKey := os.Getenv("SP_PRIVATE_KEY")
	err = ethhandler.SetArtLog(privKey, artLog.Owner, artLog.HashedData, artId)
	return err
}

// Artログをブロックチェーンから抜き出し
// func GetHashDataFromBN(artId []byte) [][]byte {
// 	reply := ethhandler.GetHashData(artId)
// 	return reply
// }

// 監査証明作成
func AuditProofGen(para *structure.Params, storage *structure.Storage) gin.HandlerFunc {
    return func(c *gin.Context) {
		artLogs := ethhandler.GetArtLogs()
		var myu *pbc.Element
		var gamma *pbc.Element
		proofs := &structure.Proofs{} 
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		g := pairing.NewG1().SetBytes(para.G)
		u := pairing.NewG1().SetBytes(para.U)
		for i := 0; i < len(artLogs); i++ {
			pubKey := pairing.NewG1().SetBytes(artLogs[i].PublicKey)
			var storageItem structure.UploadData
			for j := 0; j < len(storage.Datas); j++ {
				if reflect.DeepEqual(storage.Datas[j].ArtId, artLogs[i].Id) {
					storageItem = storage.Datas[j]
					break
				}
			}
			auditLog := ethhandler.GetAuditLog(artLogs[i].Log.LogId)
			splitedFile, _ := tool.SplitSlice(storageItem.File, len(storageItem.MetaData))
			aTable, vTable := tool.HashChallen(storageItem.SplitCount, int(auditLog.Log.Chal), auditLog.Log.K1, auditLog.Log.K2, pairing)

			var MSum *pbc.Element
			for cIndex := 0; cIndex < int(auditLog.Log.Chal); cIndex++ {
				meta := pairing.NewG1().SetBytes(storageItem.MetaData[aTable[cIndex]])
				m := pairing.NewG1().SetFromHash(splitedFile[aTable[cIndex]])
				mm := tool.GetBinaryBySHA256(m.X().String())
				log.Print("mm\n")
				log.Print(mm)
				M := pairing.NewG1().SetBytes(mm) 
				if cIndex == 0 {
					myu = pairing.NewZr().MulBig(vTable[cIndex], m.X())
					gamma = pairing.NewG1().PowZn(meta, vTable[cIndex])
				} else {
					myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[cIndex], m.X()))
					gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[cIndex]))
				}
				if cIndex == 0 {
					MSum = pairing.NewG1().PowZn(M, vTable[cIndex])
				} else {
					MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[cIndex]))
				}
			}

			uProof := pairing.NewG1().PowZn(u, myu)
			right_hand := pairing.NewG1().Mul(uProof, MSum)
			pairing_left := pairing.NewGT().Pair(gamma, g)
			pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
			log.Print("\npairing_left\n")
			log.Print(pairing_left)
			log.Print("\npairing_right\n")
			log.Print(pairing_right)

			err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
			if err != nil {
				log.Fatal("get-go-env-error")
			}
			
			privKey := os.Getenv("SP_PRIVATE_KEY")
			err = ethhandler.SetAuditProof(privKey, myu.Bytes(), gamma.Bytes(), artLogs[i].Id)
			if err != nil {
				log.Fatal("couldnt set proof")
			}
		}
		c.JSON(http.StatusOK, proofs)
	}
}
