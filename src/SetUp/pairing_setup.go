package SetUp

import (
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

func SetUp(userCount int, paramk1 int, paramk2 int) (tool.Params, tool.PrivKeys) {
	var privKeys [][]byte
	var pubKeys [][]byte
	var privKeyByte tool.PrivKeys
	var para tool.Params
	// The authority generates system parameters
	params := pbc.GenerateA(uint32(paramk1), uint32(paramk2))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	for i := 0; i < userCount; i++ {
		privKey := pairing.NewZr().Rand()
		pubKey := pairing.NewG1().MulZn(g, privKey)
		privKeys = append(privKeys, privKey.Bytes())
		pubKeys = append(pubKeys, pubKey.Bytes())
	}
	para.Pairing = params.String()
	para.G = g.Bytes()
	para.U = u.Bytes()
	para.PubKeys = pubKeys
	privKeyByte.PrivKeys = privKeys
	return para, privKeyByte
}
