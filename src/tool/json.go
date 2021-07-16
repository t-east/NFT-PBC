package tool

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/Nik-U/pbc"
)

func SaveData(file, metadata [][]byte, fileName string, inputFileString string){
	var NewStorage Storage

	NewStorage = Storage{File: file, MetaData: metadata, FileName: fileName, HashedFile: inputFileString}
	
	WriteStorage(NewStorage)
}

func ReadStorage()[]Storage{
	var storages []Storage
	rawStorage, _ := ioutil.ReadFile("../data/SP/Storage.json")
	json.Unmarshal(rawStorage, &storages)
	return storages
}

func IndexFromStorage(storages []Storage, fileTable FileIndexTable)int{ 
	for i := 0; i < len(storages); i++ {
		hashedFile, err := HashFile(storages[i].FileName)
		if err != nil {
			err.Error()
		}
		if fileTable.HashedFile == hashedFile {
			return i
		}
	}
	return 0
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

func NewFIT(data FileIndexTable) {
	var fd []FileIndexTable
	raw, err := ioutil.ReadFile("../data/BN/FileIndexTable.json")
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
	ioutil.WriteFile("../data/BN/FileIndexTable.json", content, os.ModePerm)
}

func AddFIT(data FileData, hashedFile string) {
	var fd []FileIndexTable
	raw, err := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(raw, &fd)
	for i:=0;i<len(fd);i++{
		if fd[i].HashedFile == hashedFile{
			index := i
			fd[index].UserName = append(fd[index].UserName, data.UserName)
		}
	}

	e, err := json.Marshal(fd)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := []byte(e)
	ioutil.WriteFile("../data/BN/FileIndexTable.json", content, os.ModePerm)
}

func LTToJson(data []Log) {
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := []byte(e)
	ioutil.WriteFile("../data/BN/LogTable.json", content, os.ModePerm)
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

func ParaToJson(data Params, key PrivKeys) {
	e1, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentParam := []byte(e1)
	ioutil.WriteFile("../data/BN/Para.json", contentParam, os.ModePerm)
	e2, err := json.Marshal(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	contentKey := []byte(e2)
	ioutil.WriteFile("../data/PrivKey.json", contentKey, os.ModePerm)
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

func InputPara() (Params, PrivKeys) {
	var param Params
	var privkeys PrivKeys
	rawPara, _ := ioutil.ReadFile("../data/BN/Para.json")
	rawKey, _ := ioutil.ReadFile("../data/PrivKey.json")
	json.Unmarshal(rawPara, &param)
	json.Unmarshal(rawKey, &privkeys)
	return param, privkeys
}

func InputFIT() []FileIndexTable {
	var fit []FileIndexTable
	rawFIT, _ := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(rawFIT, &fit)
	return fit
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


