package structure

import 	(
	"pairing_test/src/ethereum/ethhandler"
)

type Key struct {
	PubKey []byte `json:"public_key"`
	PrivKey []byte `json:"private_key"`
}

type User struct {
	PubKey []byte `json:"public_key"`
	PrivKey []byte `json:"private_key"`
	UserID string    `json:"user_id"`
	Address string    `json:"address"`
}

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
}

type InputFile struct {
	File []byte   `json:"file"`
	Name string `json:"name"`
}

type UploadFile struct {
	MetaData [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName string `json:"name"`
	SplitCount int `json:"split_count"`
	Owner string `json:"owner"`
}

type ArtId struct {
	Id string `json:"id"`
	Name string  `json:"name"`
}

type ArtIds struct {
	Ids []ArtId `json:"art_ids"`
}

type Chal struct {
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

func NewUser() *User {
    return &User{}
}

func NewArtIdList() *ArtIds {
    return &ArtIds{}
}

func (r *ArtIds) AddArtId(artId ArtId) {
    r.Ids = append(r.Ids, artId)
}

func File() *UploadFile {
    return &UploadFile{}
}

func NewParams() *Params {
	conn, _ := ethhandler.ConnectNetWork()
	para := ethhandler.GetPara(conn)
    result := &Params{Pairing: para.Pairing, G: para.G, U: para.U}
	return result
}

func (r *User) UserKeyGen() {
}