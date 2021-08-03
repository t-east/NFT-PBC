package SetUp

import (
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

func SetUp(userCount int, paramk1 int, paramk2 int) (tool.Params, []tool.PrivKey) {
	var pubKeys []tool.PubKey
	var privKeys []tool.PrivKey
	var para tool.Params
	// The authority generates system parameters
	params := pbc.GenerateA(uint32(paramk1), uint32(paramk2))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	for i := 0; i < userCount; i++ {
		privKey := pairing.NewZr().Rand()
		pubKey := pairing.NewG1().MulZn(g, privKey)
		pubKeyAndUser := tool.PubKey{PubKey: pubKey.Bytes(), UserID: i}
		privKeyAndUser := tool.PrivKey{PrivKey: privKey.Bytes(), UserID: i}
		privKeys = append(privKeys, privKeyAndUser)
		pubKeys = append(pubKeys, pubKeyAndUser)
	}
	para.Pairing = params.String()
	para.G = g.Bytes()
	para.U = u.Bytes()
	para.PubKeys = pubKeys
	return para, privKeys
}
