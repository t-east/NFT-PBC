package crypt

func AuditChallen(para Params) (ck, k1, k2 , err error) {
	return func(c *gin.Context) {
		pairing, _ := pbc.NewPairingFromString(para.Pairing)
		ck := rand.Intn(3) + 1
		k1 := pairing.NewZr().Rand()
		k2 := pairing.NewZr().Rand()
	}
}

// 監査検証
func AuditVerify(para, log, hashedFile) result bool {
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
	result := pairing_left.Equals(pairing_right)
	return
}