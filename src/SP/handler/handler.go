
package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
	// "encoding/binary"
	// "pairing_test/src/tool"
	// "reflect"
	"pairing_test/src/User/structure"

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
        requestBody := ArticlePostRequest{}
        c.Bind(&requestBody)

        item := structure.InputData{
			File:  requestBody.Title,
			MetaData:  requestBody.MetaData,
			UserId:  requestBody.UserId,
			ArtId:  requestBody.ArtId
        }
        storage.Add(item)
		// TODO 作品ログを記録
        c.Status(http.StatusNoContent)
    }
}

// ファイルログをブロックチェーンに保存
func SaveOutsourceData(outsourceData tool.Storage, userId int) {
	FileTable := tool.InputFIT()
	Users = append(Users, userId)
	hashedFile := tool.FileToMMData(outsourceData)
	validFile := tool.FileIndexTable{UserID: Users, HashedFile: hashedFile}
	conn, client := ethhandler.ConnectNetWork()
	auth := ethhandler.AuthUser(client)
	userAddress := ethhandler.StringToAddress("0x30b93c46e417bE95A18F5374EEAf5ca6AE457b01")
	// reply, err := conn.DeleteData(auth, validFile.HashedFile, userAddress)
	// fmt.Print("\ndelete\n")
	// if err != nil{
	// 	fmt.Print("\ndelete error\n")
	// 	fmt.Println(err)
	// }
	// fmt.Println(reply)
	
	reply, err := conn.RegisterOriginalData(auth, validFile.HashedFile, userAddress)
	if err != nil{
		fmt.Print("\nregister error\n")
		fmt.Println(err)
	}
	fmt.Println(reply)

	conn, client = ethhandler.ConnectNetWork()
	auth = ethhandler.AuthUser(client)
	userAddress = ethhandler.StringToAddress("0x30b93c46e417bE95A18F5374EEAf5ca6AE457b01")
	ethhandler.Set(auth,conn,"Hello")
	fmt.Print("\nset\n")

	ethhandler.Get(conn)
	fmt.Print("\nget\n")
	
	tool.NewFIT(validFile)
	FileTable = append(FileTable, validFile)
}

// 監査証明作成
func AuditProofGen(para *structure.Params, storage *structure.Storage, proofs *structure.OutputProof) gin.HandlerFunc {
    return func(c *gin.Context) {
		requestBody := PostChal{}
        c.Bind(&requestBody)
        challen := structure.Chal{
			C:  requestBody.C,
			K1:  requestBody.K1,
			K2:  requestBody.K2,
        }
		var myu *pbc.Element
		var gamma *pbc.Element
		var proofs structure.OutputProof
		fit := tool.InputFIT()
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for i := 0; i < len(storage); i++ {
			// ArtIndex := (storage[i].ArtIdを用いてブロックチェーンからmetadata等を取得)
			ArtIndex.HashedFile
			var proof structure.Proof
			splitedFile, _ := tool.SplitSlice(storage.Datas[i].File, len(storage.Datas[i].MetaData))
			aTable, vTable := tool.HashChallen(len(storage.Datas[i].MetaData), challen, pairing)
			for j := 0; j < challen.C; j++ {
				meta := pairing.NewG1().SetBytes(storage.Datas[i].MetaData[aTable[j]])
				m := pairing.NewG1().SetFromHash(splitedFile[aTable[j]])
				if j == 0 {
					myu = pairing.NewZr().MulBig(vTable[j], m.X())
					gamma = pairing.NewG1().PowZn(meta, vTable[j])
				} else {
					myu = pairing.NewZr().Add(myu, pairing.NewZr().MulBig(vTable[j], m.X()))
					gamma.Mul(gamma, pairing.NewG1().PowZn(meta, vTable[j]))
				}
			}
			proof := structure.Proof{
				Myu:  myu.Bytes(),
				Gamma:  gamma.Bytes(),
				ArtId: storage.Datas[i].ArtId
			}
			proofs.AddProof(proof)
		}
	}
	c.JSON(http.StatusOK, proofs)
}
