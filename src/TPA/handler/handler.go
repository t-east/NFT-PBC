
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
	reply := ethhandler.GetParam(conn)
	para.G = reply.G
	para.U = reply.U
	para.Pairing = reply.Pairing
}

func Count() gin.HandlerFunc {
	return func(c *gin.Context) {
		artLogs := ethhandler.GetArtLogs()
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, artLogs)
	}
}

func AuditChallen(para *structure.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		artLogs := ethhandler.GetArtLogs()
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		var auditIds [][]byte
		for i := 0; i < len(artLogs); i++ {
			ck, kt1, kt2 := tool.ChallenGen(3, pairing)
			logId := tool.MD5( string(ck) + string(kt1.Bytes()) + string(kt2.Bytes()) )
			auditIds = append(auditIds, logId)
			err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
			if err != nil {
				log.Fatal("get-go-env-error")
			}
			privKey := os.Getenv("TPA_PRIVATE_KEY")
			err = ethhandler.SetAuditChal(privKey, uint32(ck), kt1.Bytes(), kt2.Bytes(), logId, []byte(artLogs[i].Id))
			if err != nil {
				log.Fatal("not tpa")
			}
		}
		res, err := http.Get("http://sp:4001/proof")
		if err != nil {
			log.Fatal("[!] " + err.Error())
		} else {
			log.Print("[*] " + res.Status)
		}
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, auditIds)
	}
}


func GetAuditLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		artLogs := ethhandler.GetAuditLogs()
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, artLogs)
	}
}


// 監査検証
func AuditVerify(para *structure.Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		artLogs := ethhandler.GetArtLogs()
		for i := 0; i < len(artLogs); i++ {
			auditLog := ethhandler.GetAuditLog(artLogs[i].Log.LogId)
			aTable, vTable := tool.HashChallen(len(artLogs[i].Log.HashedFile), int(auditLog.Log.Chal), auditLog.Log.K1, auditLog.Log.K2, pairing)
			g := pairing.NewG1().SetBytes(para.G)
			pubKey := pairing.NewG1().SetBytes(artLogs[i].PublicKey)
			u := pairing.NewG1().SetBytes(para.U)
			myuT := pairing.NewZr().SetBytes(auditLog.Log.Myu)
			gammaT := pairing.NewG1().SetBytes(auditLog.Log.Gamma)
			var MSum *pbc.Element
			for c := 0; c < int(auditLog.Log.Chal); c++ {
				M := pairing.NewG1().SetBytes(artLogs[i].Log.HashedFile[aTable[c]])
				if c == 0 {
					MSum = pairing.NewG1().PowZn(M, vTable[c])
				} else {
					MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[c]))
				}
			}
			uProof := pairing.NewG1().PowZn(u, myuT)
			right_hand := pairing.NewG1().Mul(uProof, MSum)
			pairing_left := pairing.NewGT().Pair(gammaT, g)
			pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
			log.Print("\npairing_left\n")
			log.Print(pairing_left)
			log.Print("\npairing_right\n")
			log.Print(pairing_right)
			if pairing_left.Equals(pairing_right) {
				result := true
				conn, client := ethhandler.ConnectNetWork()
				err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
				if err != nil {
					log.Print("get-go-env-error")
				}
				privKey := os.Getenv("TPA_PRIVATE_KEY")
				auth := ethhandler.AuthUser(client, privKey)
				_, err = conn.SetAuditResult(auth, result, auditLog.Id )
				log.Print("Audit Verified!! ほんまに偉いよ．よく頑張った")
			}else{
				log.Print("Audit Error!!")
			}
		}
	}

}

// Artログをブロックチェーンから抜き出し
// func GetHashDataFromBN(artId []byte) contracts.IndexTableArtLogTable {
// 	reply := ethhandler.GetArtLog(artId)
// 	return reply
// }
