package SetUp

import (
	"github.com/Nik-U/pbc"
)

func SetUp() (string, []byte, []byte) {
	params := pbc.GenerateA(uint32(160), uint32(512))
	pairing := params.NewPairing()
	g := pairing.NewG1().Rand()
	u := pairing.NewG1().Rand()
	return params.String(), g.Bytes(), u.Bytes()
}
