package entities

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
}

func NewParams(pairing string, g, u []byte) *Params {
    result := &Params{Pairing: pairing, G: g, U: u}
	return result
}
