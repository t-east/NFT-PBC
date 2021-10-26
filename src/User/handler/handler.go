
package handler

import (
    "net/http"
	"fmt"
	"log"
	"io"

    "github.com/gin-gonic/gin"
	"pairing_test/src/User/structure"
	"pairing_test/src/ethereum/ethhandler"
	"pairing_test/src/tool"

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
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, para)
	}
}

// func RegisterPubKey(user User, address string) {
// 	// pubKey[address] = user.pubKeyのトランザクションを行う
// }

// ユーザ登録
func KeyGen(para *structure.Params, user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		privKey := pairing.NewZr().Rand()
		g := pairing.NewG1().SetBytes(para.G)
		pubKey := pairing.NewG1().MulZn(g, privKey)
		user.PubKey = pubKey.Bytes()
		user.PrivKey = privKey.Bytes()
		log.Print(user)
		//RegisterPubKey(user, アドレス)
        c.JSON(http.StatusOK, user)
    }
}

type EthPrivKey struct {
	PrivKey string `json:"eth_privkey"`
}

func GetAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		var requestBody EthPrivKey
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			log.Println("No extras provided")
		}
		log.Print(requestBody.PrivKey)
		address := ethhandler.GetUserAddress(requestBody.PrivKey).String()
        c.JSON(http.StatusOK, address)
    }
}

//func (r *User) SendUserId(){
	//TAにPost 
//}

type InputFile struct {
	File []byte `json:"file"`
	Name string `json:"name"`
}

// メタデータ生成
func CreateMetaData(uploadFile *structure.UploadFile, para *structure.Params, user *structure.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		u := pairing.NewG1().SetBytes(para.U)
		n := 10

		privKey := pairing.NewZr().SetBytes(user.PrivKey)

		//ファイル読み込み
		var inputFileData InputFile
		err := c.Bind(&inputFileData)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		// 署名　sk・H(m)を生成
		f, fileHeader, err := c.Request.FormFile("file")
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		defer f.Close()
        inputFile, err := io.ReadAll(f) 
		splitedFile, _ := tool.SplitSlice(inputFile, n)
		log.Print(splitedFile)
		var metaData [][]byte
		for i := 0; i < len(splitedFile); i++ {
			m := pairing.NewG1().SetFromHash(splitedFile[i])

			mm := tool.GetBinaryBySHA256(m.X().String())
			M := pairing.NewG1().SetFromHash(mm) 

			um := pairing.NewG1().PowBig(u, m.X())
			temp := pairing.NewG1().Mul(um, M)
			meta := pairing.NewG1().MulZn(temp, privKey)

			metaData = append(metaData, meta.Bytes())
		}
		uploadFile.MetaData = metaData
		uploadFile.HashedData = splitedFile
		uploadFile.Owner = user.UserID
		uploadFile.SplitCount = n
		uploadFile.FileName = fileHeader.Filename

		// // outsourceData := Chal{UserId: user.ID, MetaData: metaData, File: inputFile}
		// // SendFileToTa(outsourceData)
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, uploadFile)
	}
}

// ファイルの送信
func UploadFile(uploadFile *structure.UploadFile, para *structure.Params, user *structure.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		//ファイル読み込み
		var inputFileData InputFile
		err := c.Bind(&inputFileData)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, uploadFile)
	}
}


// 監査チャレンジ作成
// func AuditChallen(para *structure.Params, artIds *structure.ArtIds) []tool.Chal {
// 	var challens structure.Chal
// 	for i := 0; i < len(fit); i++ {
// 		pairing, _ := pbc.NewPairingFromString(sharedParams)
// 		ck, kt1, kt2 := tool.ChallenGen(len(fit[i].HashedFile), pairing)
// 		challen := tool.Chal{ID: fit[i].FileId,C: ck, K1: kt1.Bytes(), K2: kt2.Bytes()}
// 		challens = append(challens, challen)
// 	}
// 	return challens
// }

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

type UserInfo struct {
	UserID string    `json:"user_id"`
	Address string    `json:"address"`
}

func UserGet(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, user)
    }
}

func UserPost(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		requestBody := UserInfo{}
		c.Bind(&requestBody)
		user.UserID = requestBody.UserID
		user.Address = requestBody.Address 
		fmt.Print(requestBody)    
        //user.SendUserId()
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, requestBody)
    }
}

type UserRegisterInfo struct {
	UserID string    `json:"user_id"`
	Address string    `json:"address"`
	PbcPubKey string    `json:"pbc_pubkey"`
}

func Register(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		user.UserID = c.Query("user_id")
		log.Print(c.Query("user_id"))
		user.Address = c.Query("address")
		conn, client := ethhandler.ConnectNetWork()
		auth := ethhandler.AuthUser(client)
		reply, err := conn.RegisterPubKey(auth, string(user.PubKey))
		log.Print(reply)
		if err != nil {
			log.Fatal(err)
		}
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, user)
    }
}

func process(w http.ResponseWriter, r *http.Request) {

    // uploaded を参照する
    items, ok := r.MultipartForm.File["uploaded"]
    if !ok || len(items) == 0 {
        // 期待しないアップロードはエラー
        http.Error(w, "No upload files", http.StatusBadRequest)
        return
    }
    file, err := items[0].Open()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, string(data))
}