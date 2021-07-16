package TPA

import (
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

// Alice generates a keypair and signs a message
func AuditVerify(n int, params tool.Params, pubKeyByte []byte, storage tool.Storage, proofTByte [][]byte, challen tool.Chal) tool.Log {
	var log tool.Log
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	dataBlockNum := n
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)

	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes(pubKeyByte)
	u := pairing.NewG1().SetBytes(params.U)
	myuT := pairing.NewZr().SetBytes(proofTByte[0])
	gammaT := pairing.NewG1().SetBytes(proofTByte[1])
	var MSum *pbc.Element
	for i := 0; i < challen.C; i++ {
		m := pairing.NewG1().SetFromHash(storage.File[aTable[i]])
		mm := tool.GetBinaryBySHA256(m.X().String())
		M := pairing.NewG1().SetFromHash(mm)
		if i == 0 {
			MSum = pairing.NewG1().MulZn(M, vTable[i])
		} else {
			MSum.Mul(MSum, pairing.NewG1().MulZn(M, vTable[i]))
		}
	}

	uProof := pairing.NewG1().MulZn(u, myuT)
	right_hand := pairing.NewG1().Mul(uProof, MSum)
	pairing_left := pairing.NewGT().Pair(gammaT, g)
	pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
	log.Result = 0
	if pairing_left.Equals(pairing_right) {
		log.Result = 1
	}
	log.Proof = proofTByte
	log.Challen = challen
	return log

}
