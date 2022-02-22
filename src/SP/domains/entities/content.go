package entities

type ContentInput struct {
	MetaData   [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName   string   `json:"name"`
	SplitCount int      `json:"split_count"`
	Owner      string   `json:"owner"`
	ArtId      string   `json:"art_id"`
}

type ContentInStorage struct {
	Id       string   `json:"id"`
	File     []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
	FileName string   `json:"name"`
}

type Content struct {
	Id           string `gorm:"primary_key" json:"id"`
	ContentLogId string `json:"content_log_id"`
	ContentURL   string `json:"content_url"`
	FileName     string `json:"name"`
}

type ContentInBlockChain struct {
	MetaData   [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName   string   `json:"name"`
	SplitCount int      `json:"split_count"`
	Owner      string   `json:"owner"`
	ArtId      string   `json:"art_id"`
}
