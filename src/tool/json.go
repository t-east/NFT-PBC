package tool

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
	"github.com/Nik-U/pbc"
	// "net/http"
)

func SaveData(file Storage){
	var NewStorage Storage

	NewStorage = Storage{File: file.File, MetaData: file.MetaData}
	
	WriteStorage(NewStorage)
}

func ReadStorage()[]Storage{
	var storages []Storage
	rawStorage, _ := ioutil.ReadFile("../data/SP/Storage.json")
	json.Unmarshal(rawStorage, &storages)
	return storages
}

func WriteStorage(newStorage Storage){
	var storages []Storage
	raw, err := ioutil.ReadFile("../data/SP/Storage.json")
	if err != nil {
		storages = append(storages, newStorage)
	} else {
		json.Unmarshal(raw, &storages)
		storages = append(storages, newStorage)
	}
	e, err := json.Marshal(storages)
	if err != nil {
		fmt.Println(err)
	}
	content := []byte(e)
	ioutil.WriteFile("../data/SP/Storage.json", content, os.ModePerm)
	return 
}

func NewFileLog(data []byte) {
	// 作品ログを作成
}

func LocalToStorage(data Storage) {
	var fd []Storage
	raw, err := ioutil.ReadFile("../data/SP/Storage.json")
	if err != nil {
		fd = append(fd, data)
	} else {
		json.Unmarshal(raw, &fd)
		fd = append(fd, data)
	}
	e, err := json.Marshal(fd)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := []byte(e)
	ioutil.WriteFile("../data/SP/Storage.json", content, os.ModePerm)
}

func FileDataToJson(data FileData) {
	name := data.UserName
	var fd []FileData
	raw, err := ioutil.ReadFile("../data/fileData" + name + ".json")
	if err != nil {
		fd = append(fd, data)
	} else {
		json.Unmarshal(raw, &fd)
		fd = append(fd, data)
	}
	e, err := json.Marshal(fd)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := []byte(e)
	ioutil.WriteFile("../data/fileData"+name+".json", content, os.ModePerm)
}

func InputPara() (Params, []PrivKey) {
	var param Params
	var privkeys []PrivKey
	rawPara, _ := ioutil.ReadFile("../data/BN/Para.json")
	rawKey, _ := ioutil.ReadFile("../data/PrivKey.json")
	json.Unmarshal(rawPara, &param)
	json.Unmarshal(rawKey, &privkeys)
	return param, privkeys
}

func JsonCheck(name string, params Params) {
	var fd []FileData
	raw, err := ioutil.ReadFile("../data/fileData" + name + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(raw, &fd)
	pairing, _ := pbc.NewPairingFromString(params.Pairing)
	for i := 0; i < len(fd); i++ {
		filename := fd[i].FileName
		fmt.Print(filename)
		fmt.Print("\n")
		username := fd[i].UserName
		fmt.Print(username)
		fmt.Print("\n")
		dataBlockNum := fd[i].DataBlockNum
		fmt.Print(dataBlockNum)
		fmt.Print("\n")
		for j := 0; j < dataBlockNum; j++ {
			//mdata := pairing.NewG1().SetBytes(fd[i].MData[j])
			//fmt.Print(mdata)
			//fmt.Print("\n")
			mm := GetBinaryBySHA256(string(fd[i].MData[j]))
			Mdata := pairing.NewZr().SetBytes(mm)
			//Mdata := pairing.NewG1().SetBytes(fd[i].MMData[j])
			fmt.Print(Mdata)
			fmt.Print("\n")
			//metadata := pairing.NewG1().SetBytes(fd[i].MetaData[j])
			//fmt.Print(metadata)
			//fmt.Print("\n")
		}

	}

}

//TODO トランザクションAPIを触る関数を実装
// func RegisterOriginalData(){

// }

// func RegisterDedData(){

// }

// func GetFIT(ids){
	
// }

// func RegisterParam(){
	
// }

// func RegisterPubKey(pubke){

// }

// func GetParam(){
	
// }

// func GetPubKey(pubke){

// }

// func RegisterLog(pubke){

// }



