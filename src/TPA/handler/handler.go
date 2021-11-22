
package handler

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	"pairing_test/src/ethereum/ethhandler"
	"pairing_test/src/tool"
	// "encoding/binary"
	// "reflect"
	"pairing_test/src/TPA/structure"
	"fmt"
	"os"
	"github.com/joho/godotenv" 

	"github.com/Nik-U/pbc"
)

func GetPara(para *structure.Params) {
	conn, _ := ethhandler.ConnectNetWork()
	reply := ethhandler.GetPara(conn)
	para.G = reply.G
	para.U = reply.U
	para.Pairing = reply.Pairing
}

func AuditChallen(para *structure.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, _ := ethhandler.ConnectNetWork()
		reply := ethhandler.GetArtIds(conn)
		artIds := reply
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for i := 0; i < len(artIds); i++ {
			ck, kt1, kt2 := tool.ChallenGen(3, pairing)
			logId := tool.MD5( string(ck) + string(kt1.Bytes()) + string(kt2.Bytes()) )
			conn, client := ethhandler.ConnectNetWork()
			err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
			if err != nil {
				log.Fatal("get-go-env-error")
			}
			privKey := os.Getenv("TPA_PRIVATE_KEY")
			auth := ethhandler.AuthUser(client, privKey)
			_, err = conn.SetLogId(auth, uint32(ck), kt1.Bytes(), kt2.Bytes(), []byte(artIds[i]), logId)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, "OK")
	}
}

// 監査検証
func AuditVerify(para *structure.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		logs := ethhandler.GetLog()
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for p := 0; p < len(logs); p++ {
			hashedFile := ethhandler.GetHashData(logs[p].Log.ArtId)
			aTable, vTable := tool.HashChallen(len(hashedFile), int(logs[p].Log.Chal), logs[p].Log.Myu, logs[p].Log.Gamma, pairing)
			g := pairing.NewG1().SetBytes(para.G)
			pubKey := pairing.NewG1().SetBytes(ethhandler.GetPubkey(logs[p].Log.ArtId))
			u := pairing.NewG1().SetBytes(para.U)
			myuT := pairing.NewZr().SetBytes(logs[p].Log.Myu)
			gammaT := pairing.NewG1().SetBytes(logs[p].Log.Gamma)
			var MSum *pbc.Element
			for i := 0; i < int(logs[p].Log.Chal); i++ {
				M := pairing.NewG1().SetBytes(hashedFile[aTable[i]])
				if i == 0 {
					MSum = pairing.NewG1().PowZn(M, vTable[i])
				} else {
					MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[i]))
				}
			}
			uProof := pairing.NewG1().PowZn(u, myuT)
			right_hand := pairing.NewG1().Mul(uProof, MSum)
			pairing_left := pairing.NewGT().Pair(gammaT, g)
			pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
			if pairing_left.Equals(pairing_right) {
				result := true
				conn, client := ethhandler.ConnectNetWork()
				err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
				if err != nil {
					log.Print("get-go-env-error")
				}
				privKey := os.Getenv("TPA_PRIVATE_KEY")
				auth := ethhandler.AuthUser(client, privKey)
				_, err = conn.SetAuditResult(auth, result, logs[p].LogId )
				}else{
					log.Print("Audit Error!!")
				}
		}
	}

}

// Artログをブロックチェーンから抜き出し
func GetHashDataFromBN(artId []byte) [][]byte {
	reply := ethhandler.GetHashData(artId)
	return reply
}
