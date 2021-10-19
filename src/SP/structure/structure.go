package structure

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
}

type InputData struct {
	File []byte   `json:"file"`
	MetaData       [][]byte   `json:"meta_data"`
	UserId       string   `json:"user_id"`
	ArtId string `json:"art_id"`
}

type Storage struct {
	Datas []InputData `json:"datas"`
}

type Chal struct {
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

type Proof struct {
	Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
}

type OutputProof struct {
	Proofs []Proof
}

func Datas() *OutsourceDatas {
    return &OutsourceDatas{}
}

func NewStorage() *Storage {
    return &Storage{}
}

func (r *Storage) AddStorage(data InputData) {
    r.Datas = append(r.Datas, data)
}

func (r *OutputProof) AddProof(proof Proof) {
    r.Proofs = append(r.Proofs, proof)
}

func NewParams() *Params {
    return &Params{}
}

func (r *User) UserKeyGen() {
}