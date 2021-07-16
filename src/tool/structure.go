package tool

type FileIndexTable struct {
	UserName   []string `json:"user_name"`
	HashedFile string   `json:"hashed_file"`
}

type Storage struct {
	FileName   string   `json:"file_name"`
	HashedFile string   `json:"hashed_file"`
	File       [][]byte `json:"file"`
	MetaData   [][]byte `json:"meta_data"`
}

type Chal struct {
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
}

type DataList struct {
	Name     string `json:"name"`
	DataName string `json:"data_name"`
}

type Log struct {
	Challen Chal     `json:"challenge"`
	Result  int      `json:"result"`
	Proof   [][]byte `json:"proof"`
}

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
	PubKeys [][]byte `json:"public_keys"`
}

type PrivKeys struct {
	PrivKeys [][]byte `json:"private_keys"`
}

type FileData struct {
	UserName     string   `json:"user_name"`
	FileName     string   `json:"file_name"`
	MetaData     [][]byte `json:"meta_data"`
	MMData       [][]byte `json:"hashed_m_data"`
	MData        [][]byte `json:"m_data"`
	DataBlockNum int      `json:"data_block_num"`
	SplitedData  [][]byte `json:"splited_data"`
}

func DataListGen() []DataList {
	dataList := []DataList{}
	dataList = append(dataList, DataList{"A", "inputData/linux_logo.jpg"})
	dataList = append(dataList, DataList{"B", "inputData/testfile"})
	dataList = append(dataList, DataList{"B", "inputData/linux_logo.jpg"})
	dataList = append(dataList, DataList{"C", "inputData/testfile"})
	return dataList
}
