package ethereum

type Param struct {
	Pairing string
	G       string
	U       string
}

func GetParam() (Param, error) {
	return Param{}, nil
}
