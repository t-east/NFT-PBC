package tool

type FileIndexTable struct {
	UserID     []int    `json:"user_id"`
	HashedFile [][]byte `json:"hashed_file"`
	FileId     int      `json:"file_id"`
}

type Storage struct {
	Id       int      `json:"id"`
	File     []byte   `json:"file"`
	MetaData [][]byte `json:"meta_data"`
}

type Chal struct {
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
	ID int    `json:"id"`
}

type ProofT struct {
	ID    int    `json:"id"`
	Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
}

type DataList struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	DataName string `json:"data_name"`
}

type Log struct {
	Challen Chal   `json:"challenge"`
	Result  int    `json:"result"`
	Proof   ProofT `json:"proof"`
	FileId  int    `json:"file_id"`
}

type PubKey struct {
	PubKey []byte `json:"public_keys"`
	UserID int    `json:"user_id"`
}

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
	PubKeys []PubKey `json:"public_keys"`
}

type UserTable struct {
	User []string `json:"user_table"`
}

type PrivKey struct {
	PrivKey []byte `json:"private_key"`
	UserID  int    `json:"user_id"`
}

type FileData struct {
	UserName     string   `json:"user_name"`
	UserID       int      `json:"user_id"`
	FileName     string   `json:"file_name"`
	MetaData     [][]byte `json:"meta_data"`
	MMData       [][]byte `json:"hashed_m_data"`
	MData        [][]byte `json:"m_data"`
	DataBlockNum int      `json:"data_block_num"`
	SplitedData  [][]byte `json:"splited_data"`
}

type OutsourceData struct {
	UserName string   `json:"user_name"`
	UserID   int      `json:"user_id"`
	FileName string   `json:"file_name"`
	MetaData [][]byte `json:"meta_data"`
	File     []byte   `json:"file"`
}

func DataListGen() []DataList {
	dataList := []DataList{}
	dataList = append(dataList, DataList{Name: "A", ID: 0, DataName: "inputData/linux_logo.jpg"})
	dataList = append(dataList, DataList{Name: "C", ID: 2, DataName: "inputData/testfile"})
	dataList = append(dataList, DataList{Name: "A", ID: 0, DataName: "inputData/testfile"})
	dataList = append(dataList, DataList{Name: "B", ID: 1, DataName: "inputData/linux_logo.jpg"})
	return dataList
}

func Contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
