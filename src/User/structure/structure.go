package structure

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

type OutsourceData struct {
	File []byte   `json:"file"`
	MetaData       [][]byte   `json:"meta_data"`
	UserId       string   `json:"user_id"`
}

type OutsourceDatas struct {
	Data []OutsourceData `json:"datas"`
}

type ArtId struct {
	Id string `json:"id"`
	Name string  `json:"name"`
}

type ArtIds struct {
	Ids []ArtId `json:"art_ids"`
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

func Datas() *OutsourceDatas {
    return &OutsourceDatas{}
}

func NewData() *OutsourceData {
    return &OutsourceData{}
}

func NewParams() *Params {
    return &Params{}
}

func (r *User) UserKeyGen() {
}