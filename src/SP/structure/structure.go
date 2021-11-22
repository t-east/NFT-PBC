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

type UploadData struct {
	File []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
	SplitCount int `json:"split_count"`
	FileName string `json:"name"`
	Owner string `json:"owner"`
	ArtId []byte `json:"art_id"`
}

type Storage struct {
	Datas []UploadData `json:"datas"`
}

type ArtLog struct {
	HashedData [][]byte `json:"hashed_data"`
	Owner string `json:"owner"`
}

type UserRequest struct {
	ArtLog ArtLog `json:"art_log"`
	UploadData UploadData `json:"upload_data"`
	ArtId string `json:"art_id"`
}

type UploadFile struct {
	File []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName string `json:"name"`
	SplitCount int `json:"split_count"`
	Owner string `json:"owner"`
	ArtId string `json:"art_id"`
}

type HashedResponseToUser struct {
	ArtId string  `json:"art_id"`
	HashedFile [][]byte `json:"hashed_file"`
}

type Chal struct {
	ArtId string `json:"art_id`
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

type Proof struct {
	Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
	ArtId string  `json:"art_id"`
}

type Proofs struct {
	Proofs   []Proof `json:"proofs"`
}

func NewStorage() *Storage {
    return &Storage{}
}

func (r *Storage) AddStorage(data UploadData) {
    r.Datas = append(r.Datas, data)
}

func (p *Proofs) AddProofs(proof Proof) {
    p.Proofs = append(p.Proofs, proof)
}

func NewParams() *Params {
    return &Params{}
}
