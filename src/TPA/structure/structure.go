package structure

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
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