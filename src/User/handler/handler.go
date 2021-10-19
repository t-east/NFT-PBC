
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

// func RegisterPubKey(user User, address string) {
// 	// pubKey[address] = user.pubKeyのトランザクションを行う
// }

// ユーザ登録
func KeyGenPost(para *structure.Params, user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		privKey := pairing.NewZr().Rand()
		g := pairing.NewG1().SetBytes(para.G)
		pubKey := pairing.NewG1().MulZn(g, privKey)
		user.PubKey = pubKey.Bytes()
		user.PrivKey = privKey.Bytes()
		//RegisterPubKey(user, アドレス)
        c.JSON(http.StatusOK, user)
    }
}

//func (r *User) SendUserId(){
	//TAにPost 
//}

// メタデータ生成
// func (r *User) Upload(inputData InputData) OutsourceData {
// 	para := GetPara()
// 	pairing, _ := pbc.NewPairingFromString(para.Pairing)
// 	u := pairing.NewG1().SetBytes(para.U)
// 	n := 5

// 	privKey := pairing.NewZr().SetBytes(user.PrivKey)

// 	// 署名　sk・H(m)を生成
// 	inputFile := tool.ReadBinaryFile(inputData.DataPath, binary.BigEndian)
// 	splitedFile, _ := tool.SplitSlice(inputFile, n)
// 	var metaData [][]byte
// 	for i := 0; i < len(splitedFile); i++ {
// 		m := pairing.NewG1().SetFromHash(splitedFile[i])

// 		mm := tool.GetBinaryBySHA256(m.X().String())
// 		M := pairing.NewG1().SetFromHash(mm) 

// 		um := pairing.NewG1().PowBig(u, m.X())
// 		temp := pairing.NewG1().Mul(um, M)
// 		meta := pairing.NewG1().MulZn(temp, privKey)

// 		metaData = append(metaData, meta.Bytes())
// 	}
// 	outsourceData := Chal{UserId: user.ID, MetaData: metaData, File: inputFile}
// 	// SendFileToTa(outsourceData)
// 	return outsourceData
// }

// 監査チャレンジ作成
func AuditChallen(para *structure.Params, artIds *structure.ArtIds) []tool.Chal {
	var challens []tool.Chal
	for i := 0; i < len(fit); i++ {
		pairing, _ := pbc.NewPairingFromString(sharedParams)
		ck, kt1, kt2 := tool.ChallenGen(len(fit[i].HashedFile), pairing)
		challen := tool.Chal{ID: fit[i].FileId,C: ck, K1: kt1.Bytes(), K2: kt2.Bytes()}
		challens = append(challens, challen)
	}
	return challens
}

// func CheckDep(outSourceData tool.Storage) int {
// 	found := 0
// 	fit := tool.InputFIT()
// 	for i := 0; i < len(fit); i++ {
// 		hashedFile := tool.FileToMMData(outSourceData)
// 		if reflect.DeepEqual(fit[i].HashedFile, hashedFile) {
// 			found = 1
// 		}
// 	}
// 	return found
// }

// //　監査ログ確認
// func CheckLog(user int, params Params) int {
// 	result := 0
// 	fit := tool.InputFIT()
// 	var linkIdTable []int

// 	//FIT
// 	for i := 0; i < len(fit); i++ {
// 		if tool.Contains(fit[i].UserID, user) {
// 			linkIdTable = append(linkIdTable, fit[i].FileId)
// 		}
// 	}
// 	//Para
// 	pairing, _ := pbc.NewPairingFromString(params.Pairing)
// 	u := pairing.NewG1().SetBytes(params.U)
// 	g := pairing.NewG1().SetBytes(params.G)
// 	logs := tool.ReadLogs()
// 	var challen tool.Chal
// 	var hashedFile [][]byte
// 	var pubKey *pbc.Element
// 	for j := 0; j < len(logs); j++ {
// 		if tool.Contains(linkIdTable, logs[j].FileId) {
// 			proof := pairing.NewZr().SetBytes(logs[j].Proof.Myu)
// 			uProof := pairing.NewG1().MulZn(u, proof)
// 			challen.C = logs[j].Challen.C
// 			challen.K1 = logs[j].Challen.K1
// 			challen.K2 = logs[j].Challen.K2
// 			gammaT := pairing.NewG1().SetBytes(logs[j].Proof.Gamma)
// 			aTable, vTable := tool.HashChallen(10, challen, pairing)
// 			for _, v := range fit {
// 				if v.FileId == logs[j].FileId {
// 					hashedFile = v.HashedFile
// 					pubKey = pairing.NewG1().SetBytes(params.PubKeys[v.UserID[0]].PubKey)
// 					break
// 				}
// 			}
// 			x := uProof
// 			for i := 0; i < challen.C; i++ {
// 				M := pairing.NewG1().SetBytes(hashedFile[aTable[i]])
// 				Mv := pairing.NewG1().MulZn(M, vTable[i])
// 				x = pairing.NewG1().Mul(x, Mv)
// 			}
// 			pairing_left := pairing.NewGT().Pair(gammaT, g)
// 			pairing_right := pairing.NewGT().Pair(x, pubKey)
// 			if pairing_left.Equals(pairing_right) {
// 				result = 1
// 			}
// 		}
// 	}
// 	return result
// }

type User struct {
	PubKey []byte `json:"public_key"`
	PrivKey []byte `json:"private_key"`
	UserID string    `json:"user_id"`
	Address string    `json:"address"`
}

func UserGet(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, user)
    }
}

func UserPost(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		requestBody := structure.User{}
        c.Bind(&requestBody)
        //user.SendUserId()
        c.Status(http.StatusNoContent)
    }
}

func FilePost(datas *structure.OutsourceDatas) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := structure.InputFile{}
        c.Bind(&requestBody)
        c.Status(http.StatusNoContent)
    }
}