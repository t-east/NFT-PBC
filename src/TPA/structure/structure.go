package structure

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
}

type Chal struct {
	ArtId []byte `json:"art_id`
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

type Proof struct {
	Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
	ArtId []byte  `json:"art_id"`
}

type AuditResult struct {
	Chal Chal `json:"challenge"`
	Proof Proof `json:"proof"`
	Result int `json:"result"`
}

type ArtId struct {
	Id string `json:"id"`
	Name string  `json:"name"`
}

type ArtIds struct {
	Ids [][]byte `json:"art_ids"`
}

type Log struct {
    Result bool `json:"result"`
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
    Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
	ArtId string  `json:"art_id"`
	LogId string  `json:"log_id"`
  }

func NewParams() *Params {
    return &Params{}
}