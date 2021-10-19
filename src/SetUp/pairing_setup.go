package SetUp

import (
	"pairing_test/src/tool"

	"github.com/Nik-U/pbc"
)

func SetUp(userCount int, paramk1 int, paramk2 int) (tool.Params, []tool.PrivKey) {
	var para tool.Params
	// The authority generates system parameters
	params := pbc.GenerateA(uint32(paramk1), uint32(paramk2))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	para.Pairing = params.String()
	para.G = g.Bytes()
	para.U = u.Bytes()
	// TODO:ブロックチェーンにParaを保存
	return para
}

func RegisterUser(id string, para tool.Params) tool.UserKey {
	var userKey tool.UserKey
	pairing, _ := pbc.NewPairingFromString(para.Pairing)
	privKey := pairing.NewZr().Rand()
	g := pairing.NewG1().SetBytes(para.G)
	pubKey := pairing.NewG1().MulZn(g, privKey)
	userKey.ID = id
	userKey.PubKey = pubKey.Bytes()
	userKey.PrivKey = privKey.Bytes()
	return userKey
}
