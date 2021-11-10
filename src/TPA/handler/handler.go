
package handler

import (
    "net/http"
	"log"
    "github.com/gin-gonic/gin"
	// "encoding/binary"
	// "reflect"
	"pairing_test/src/TPA/structure"

	"github.com/Nik-U/pbc"
)

func GetPara(para *structure.Params) {
	conn, _ := ethhandler.ConnectNetWork()
	reply := ethhandler.GetPara(conn)
	para.G = reply.G
	para.U = reply.U
	para.Pairing = reply.Pairing
}

func PostChallens(challens *[]structure.Chal) {
	return func(c *gin.Context) {
		challensInput := []structure.Chal{}
		c.Bind(&challensInput)
		challens = challensInput
		log.Print(challens)
		c.Header("Access-Control-Allow-Origin", "*")
        c.JSON(http.StatusOK, challens)
	}
}

// 監査検証
func AuditVerify(para *structure.Params, challens *[]structure.Chal) ([]tool.Log, int){
	return func(c *gin.Context) {
		proofs := []structure.Proof{}
		logs := []structure.Log{}
		c.Bind(&proofs)
		errors := 1
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		for p := 0; p < len(proofs); p++ {
			var chal structure.Chal
			for cIndex := 0; cIndex < len(challens); p++ {
				if challens[cIndex].ArtId == proofs[p].ArtId {
					chal = challens[cIndex]
				}
			}
			hashedFile := GetHashDataFromBN(proofs[p].ArtId)
			aTable, vTable := tool.HashChallen(len(hashedFile), chal, pairing)
			g := pairing.NewG1().SetBytes(para.G)
			pubKey := pairing.NewG1().SetBytes(para.PubKeys[fit[p].UserID[0]].PubKey)
			u := pairing.NewG1().SetBytes(para.U)
			myuT := pairing.NewZr().SetBytes(proofs[p].Myu)
			gammaT := pairing.NewG1().SetBytes(proofs[p].Gamma)
			var MSum *pbc.Element
			for i := 0; i < chal.C; i++ {
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
			var log structure.Log
			if pairing_left.Equals(pairing_right) {
				log.Result = 1
			}else{
				errors = 0
				fmt.Print(errors)
			}
			log.Myu = proofs[p].Myu
			log.Gamma = proofs[p].Gamma
			log.C = chal.C
			log.K1 = chal.K1
			log.K2 = chal.K2
			log.ArtId = proofs[p].ArtId
			log.LogId := sha1.Sum([]byte( string(log.Myu) + string(log.Gamma) + string(log.C) + string(log.K1) + string(log.K2) + log.ArtId )
			logs = append(logs, log)
		}
	}

}

// Artログをブロックチェーンから抜き出し
func GetHashDataFromBN(artId string) [][]byte {
	conn, _ := ethhandler.ConnectNetWork()	
	reply := ethhandler.GetHashData(conn, artId)
	return reply
}
