package TPA

import (
	"pairing_test/src/tool"
	"fmt"
	"github.com/Nik-U/pbc"
)

// Alice generates a keypair and signs a message
func AuditVerify(params tool.Params, proofTByte []tool.ProofT, challen []tool.Chal) ([]tool.Log, int){
	var log tool.Log
	fit := tool.InputFIT()
	var logs []tool.Log
	errors := 1
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	for l := 0; l < len(fit); l++ {
		aTable, vTable := tool.HashChallen(len(fit[l].HashedFile), challen[l], pairing)
		g := pairing.NewG1().SetBytes(params.G)
		pubKey := pairing.NewG1().SetBytes(params.PubKeys[fit[l].UserID[0]].PubKey)
		u := pairing.NewG1().SetBytes(params.U)
		myuT := pairing.NewZr().SetBytes(proofTByte[l].Myu)
		gammaT := pairing.NewG1().SetBytes(proofTByte[l].Gamma)
		var MSum *pbc.Element
		for i := 0; i < challen[l].C; i++ {
			M := pairing.NewG1().SetBytes(fit[l].HashedFile[aTable[i]])
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
		log.Result = 0
		if pairing_left.Equals(pairing_right) {
			log.Result = 1
		}else{
			errors = 0
			fmt.Print(errors)
		}
		log.Proof = proofTByte[l]
		log.Challen = challen[l]
		log.FileId = fit[l].FileId
		logs = append(logs, log)
	}
	return logs, errors

}
