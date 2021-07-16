package SP

import (
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

func DedupChallen(sharedParams string, n int) tool.Chal {
	pairing, _ := pbc.NewPairingFromString(sharedParams)
	ck, ks1, ks2 := tool.ChallenGen(n, pairing)
	challen := tool.Chal{C: ck, K1: ks1.Bytes(), K2: ks2.Bytes()}
	return challen
}

// Alice generates a keypair and signs a message
func DedupVerify(fileData tool.FileData, params tool.Params, pubKeyByte []byte, proofs []byte, challen tool.Chal) int {
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	dataBlockNum := fileData.DataBlockNum
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)
	g := pairing.NewG1().SetBytes(params.G)
	pubKey := pairing.NewG1().SetBytes(pubKeyByte)
	u := pairing.NewG1().SetBytes(params.U)
	proof := pairing.NewZr().SetBytes(proofs)

	var MSum *pbc.Element
	var metaSum *pbc.Element

	for i:=0;i<challen.C;i++{
		M := pairing.NewG1().SetBytes(fileData.MMData[aTable[i]])
		meta := pairing.NewG1().SetBytes(fileData.MetaData[aTable[i]])

		if i == 0 {
			metaSum = pairing.NewG1().PowZn(meta, vTable[i])
			MSum = pairing.NewG1().PowZn(M,vTable[i])
		} else {
			MSum.Mul(MSum, pairing.NewG1().PowZn(M, vTable[i]))
			metaSum.Mul(metaSum, pairing.NewG1().PowZn(meta, vTable[i]))
		}
	}

	uProof := pairing.NewG1().PowZn(u, proof)
	right_hand := pairing.NewG1().Mul(uProof, MSum)
	pairing_left := pairing.NewGT().Pair(metaSum, g)
	pairing_right := pairing.NewGT().Pair(right_hand, pubKey)
	DetupVerifyResult := 0
	if pairing_left.Equals(pairing_right) {
		DetupVerifyResult = 1
	}
	return DetupVerifyResult
}

func AuditProofGen(n int, storage tool.Storage, params tool.Params, challen tool.Chal) [][]byte {
	var myuT *pbc.Element
	var gammaT *pbc.Element
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	dataBlockNum := n
	aTable, vTable := tool.HashChallen(dataBlockNum, challen, pairing)
	for i := 0; i < challen.C; i++ {
		meta := pairing.NewG1().SetBytes(storage.MetaData[aTable[i]])
		m := pairing.NewG1().SetFromHash(storage.File[aTable[i]])
		if i == 0 {
			myuT = pairing.NewZr().MulBig(vTable[i],m.X())
			gammaT = pairing.NewG1().PowZn(meta, vTable[i])
		} else {
			myuT.Add(myuT, pairing.NewZr().MulBig(vTable[i],m.X()))
			gammaT.Mul(gammaT, pairing.NewG1().PowZn(meta, vTable[i]))
		}
	}
	var proofT [][]byte
	proofT = append(proofT, myuT.Bytes())
	proofT = append(proofT, gammaT.Bytes())
	return proofT
}

