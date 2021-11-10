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

type UploadFileFromUser struct {
	MetaData [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName string `json:"name"`
	SplitCount int `json:"split_count"`
	Owner string `json:"owner"`
	ArtId string `json:"art_id"`
}

type ArtLog struct {
	HashedData [][]byte `json:"hashed_data"`
	Owner string `json:"owner"`
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

type ArtId struct {
	Id string `json:"id"`
	Name string  `json:"name"`
}

type ArtIds struct {
	Ids []ArtId `json:"art_ids"`
}

type Chal struct {
	ArtId string `json:"art_id`
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