
package handler

import (
    "net/http"
	"fmt"
	"log"
	"io"
	"os"
	"encoding/json"
	"bytes"
	"github.com/joho/godotenv" 

    "github.com/gin-gonic/gin"
	"pairing_test/src/User/structure"
	"pairing_test/src/ethereum/ethhandler"
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

func GetPara(para *structure.Params) gin.HandlerFunc {
    return func(c *gin.Context) {
		conn, _ := ethhandler.ConnectNetWork()
		reply := ethhandler.GetPara(conn)
		para.G = reply.G
		para.U = reply.U
		para.Pairing = reply.Pairing
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusOK, para)
	}
}

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
		address := ethhandler.GetUserAddress(requestBody.PrivKey).String()
        c.JSON(http.StatusOK, address)
    }
}

func UserGet(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, user)
    }
}

type UserInfo struct {
	UserID string    `json:"user_id"`
	Address string    `json:"address"`
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

func Register(user *structure.User) gin.HandlerFunc {
    return func(c *gin.Context) {
		user.UserID = c.Query("user_id")
		user.Address = c.Query("address")
		conn, client := ethhandler.ConnectNetWork()
		err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal("get-go-env-error")
		}
		privKey := os.Getenv("USER_PRIVATE_KEY")
		auth := ethhandler.AuthUser(client, privKey)
		_, err = conn.RegisterPubKey(auth, string(user.PubKey))
		if err != nil {
			log.Fatal(err)
		}
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, user)
    }
}

type InputFile struct {
	File []byte `json:"file"`
	Name string `json:"name"`
}

// メタデータ生成
func CreateMetaData(uploadFile *structure.UploadFile, para *structure.Params, user *structure.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		u := pairing.NewG1().SetBytes(para.U)
		n := 3

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
		var metaData [][]byte
		metaToHash := ""
		for i := 0; i < len(splitedFile); i++ {
			m := pairing.NewG1().SetFromHash(splitedFile[i])

			mm := tool.GetBinaryBySHA256(m.X().String())
			M := pairing.NewG1().SetFromHash(mm) 

			um := pairing.NewG1().PowBig(u, m.X())
			temp := pairing.NewG1().Mul(um, M)
			meta := pairing.NewG1().MulZn(temp, privKey)

			metaData = append(metaData, meta.Bytes())
			metaToHash = metaToHash + meta.String()
		}

		uploadFile.ArtId = tool.MD5(metaToHash)
		uploadFile.MetaData = metaData
		uploadFile.HashedData = splitedFile
		uploadFile.Owner = user.UserID
		uploadFile.SplitCount = n
		uploadFile.FileName = fileHeader.Filename
		uploadFile.File = inputFile
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, uploadFile)
	}
}

// ファイルの送信
func UploadFile(artIds *structure.ArtIds, uploadFile *structure.UploadFile, para *structure.Params, user *structure.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqJson := new(bytes.Buffer)
		log.Print(uploadFile.HashedData)
		json.NewEncoder(reqJson).Encode(uploadFile)
		res, err := http.Post("http://sp:4001/art", "application/json", reqJson)
		if err != nil {
			log.Fatal("[!] " + err.Error())
		} else {
			log.Print("[*] " + res.Status)
		}
		var artLog structure.ArtLog
		artIds.Ids = append(artIds.Ids, structure.ArtId{Id: uploadFile.ArtId, Name: uploadFile.FileName})
		err = json.NewDecoder(res.Body).Decode(&artLog)
		c.Header("Access-Control-Allow-Origin", "*")
		log.Print(artLog)
		log.Print(artIds)
        c.JSON(http.StatusOK, artLog)

	}
}

//　ファイルをストレージに保存
func ArtGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		artId := c.Query("id")
		conn, _ := ethhandler.ConnectNetWork()	
		hashedFile := ethhandler.GetHashData(conn, artId)
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, hashedFile)
	}
}

func AuditChallen(para *structure.Params, artIds *structure.ArtIds) gin.HandlerFunc {
	return func(c *gin.Context) {
		var challens []structure.Chal
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for i := 0; i < len(artIds.Ids); i++ {
			ck, kt1, kt2 := tool.ChallenGen(3, pairing)
			challen := structure.Chal{ArtId: artIds.Ids[i].Id ,C: ck, K1: kt1.Bytes(), K2: kt2.Bytes()}
			challens = append(challens, challen)
		}
		reqJson := new(bytes.Buffer)
		json.NewEncoder(reqJson).Encode(challens)
		var resSP *http.Response
		resTPA, err := http.Post("http://tpa:4002/chal", "application/json", reqJson)
		if err != nil{
			log.Fatal("tpa has not get chal")
		} else {
			log.Print(resTPA)
			resSP, err = http.Post("http://sp:4001/proof", "application/json", reqJson)
			if err != nil{
				log.Fatal("sp has not get chal")
			}
			log.Print(resSP)
		}
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, resSP.Body)
	}
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