
package handler

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	// "encoding/binary"
	// "pairing_test/src/tool"
	// "reflect"
	"pairing_test/src/SP/structure"
	"pairing_test/src/ethereum/ethhandler"

	"github.com/Nik-U/pbc"
)

func GetPara(para *structure.Params) gin.HandlerFunc {
    return func(c *gin.Context) {
		params := pbc.GenerateA(uint32(160), uint32(512))
		pairing := params.NewPairing()
		g := pairing.NewG1().Rand()
		u := pairing.NewG1().Rand()
		para.Pairing = params.String()
		para.G = g.Bytes()
		para.U = u.Bytes()
		c.JSON(http.StatusOK, para)
	}
}

//　ファイルをストレージに保存
func ArtPost(storage *structure.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
        requestBody := structure.UserRequest{}
        c.Bind(&requestBody)

        uploadItem := structure.UploadData{
			File: requestBody.UploadData.File,
			MetaData: requestBody.UploadData.MetaData,
			FileName: requestBody.UploadData.FileName,
			Owner: requestBody.UploadData.Owner}

		artLog := structure.ArtLog{
			HashedData:  requestBody.ArtLog.HashedData,
			Owner:  requestBody.ArtLog.Owner,}
		
		SaveOutsourceData(artLog)
        storage.AddStorage(uploadItem)
		// TODO 作品ログを記録
        c.Status(http.StatusNoContent)
    }
}

// ファイルログをブロックチェーンに保存
func SaveOutsourceData(artLog structure.ArtLog) {
	conn, client := ethhandler.ConnectNetWork()
	auth := ethhandler.AuthUser(client)
	userAddress := ethhandler.StringToAddress(artLog.Owner)
	
	reply, err := conn.RegisterOriginalData(auth, artLog.HashedData, userAddress)
	if err != nil{
		log.Print("\nregister error\n")
		log.Fatal(err)
	}
	log.Print(reply)
}

// // 監査証明作成
// func AuditProofGen(para *structure.Params, storage *structure.Storage, proofs *structure.OutputProof) gin.HandlerFunc {
//     return func(c *gin.Context) {
// 		requestBody := PostChal{}
//         c.Bind(&requestBody)
//         challen := structure.Chal{
// 			C:  requestBody.C,
// 			K1:  requestBody.K1,
// 			K2:  requestBody.K2,
//         }
// 		var myu *pbc.Element
// 		var gamma *pbc.Element
// 		var proofs structure.OutputProof
// 		fit := tool.InputFIT()
// 		pairing, _ := pbc.NewPairingFromString(para.Pairing)
// 		for i := 0; i < len(storage); i++ {
// 			// ArtIndex := (storage[i].ArtIdを用いてブロックチェーンからmetadata等を取得)
// 			ArtIndex.HashedFile
// 			splitedFile, _ := tool.SplitSlice(storage.Datas[i].File, len(storage.Datas[i].MetaData))
// 			aTable, vTable := tool.HashChallen(len(storage.Datas[i].MetaData), challen, pairing)
// 			for j := 0; j < challen.C; j++ {
// 				meta := pairing.NewG1().SetBytes(storage.Datas[i].MetaData[aTable[j]])
// 				m := pairing.NewG1().SetFromHash(splitedFile[aTable[j]])
// 				if j == 0 {
// 					myu = pairing.NewZr().MulBig(vTable[j], m.X())
// 					gamma = pairing.NewG1().PowZn(meta, vTable[j])
// 				} else {
// 					myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[j], m.X()))
// 					gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[j]))
// 				}
// 			}
// 			proof := structure.Proof{
// 				Myu:  myu.Bytes(),
// 				Gamma:  gamma.Bytes(),
// 				ArtId: storage.Datas[i].ArtId}
// 			proofs.AddProof(proof)
// 		}
// 	}
// 	c.JSON(http.StatusOK, proofs)
// }
