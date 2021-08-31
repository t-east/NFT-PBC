package tool

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Nik-U/pbc"
	// "net/http"
)

func MakeDb(){
	db := gormConnect()
	defer db.Close()

  db.AutoMigrate(&Log{})
}
func gormConnect() *gorm.DB {
  DBMS   := "mysql"
	DBUser := "root"
	DBPass := "root"
	DBProtocol := "tcp(192.168.0.1:3306)"
	DBName := "test"
  connectTemplate := "%s:%s@%s/%s"
  connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
  db,err := gorm.Open(DBMS, connect)

  if err != nil {
    panic(err.Error())
  }
  return db
}

func SaveData(file Storage){
	var NewStorage Storage

	NewStorage = Storage{File: file.File, MetaData: file.MetaData}
	
	WriteStorage(NewStorage)
}

func GetFITInfoFromSP(storage Storage)int{
	var fit []FileIndexTable
	rawFIT, _ := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(rawFIT, &fit)
	for i, v := range fit {
		hashedFile := FileToMMData(storage)
		if reflect.DeepEqual(hashedFile, v.HashedFile) {
			return i
		}
	}
	return 0
}

func GetFITInfoFromUser(MData [][]byte)int{
	var fit []FileIndexTable
	rawFIT, _ := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(rawFIT, &fit)
	for i, v := range fit {
		if reflect.DeepEqual(MData, v.HashedFile) {
			return i
		}
	}
	return 0
}


func ReadStorage()[]Storage{
	var storages []Storage
	rawStorage, _ := ioutil.ReadFile("../data/SP/Storage.json")
	json.Unmarshal(rawStorage, &storages)
	return storages
}

func SesrchStorage(user int)[]Storage{
	var storages []Storage
	var fit []FileIndexTable
	var resultTable []Storage
	rawStorage, _ := ioutil.ReadFile("../data/SP/Storage.json")
	rawFIT, _ := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(rawStorage, &storages)
	json.Unmarshal(rawFIT, &fit)
	for f:=0; f < len(fit); f++{
		for _, v := range fit[f].UserID {
			if user == v {
				for _, k := range storages {
					hashedFile := FileToMMData(k)
					if reflect.DeepEqual(fit[f].HashedFile, hashedFile){
						resultTable = append(resultTable, k)
					}
				}
			}
		}
	}
	return resultTable
}

func PullFile() {

}

func ReadLogs()[]Log{
	var logs []Log
	rawLog, _ := ioutil.ReadFile("../data/BN/LogTable.json")
	json.Unmarshal(rawLog, &logs)
	return logs
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
		data.FileId = 0
		fd = append(fd, data)
	} else {
		json.Unmarshal(raw, &fd)
		data.FileId = len(fd)
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

func AddFIT(data Storage,userId int) {
	var fd []FileIndexTable
	raw, err := ioutil.ReadFile("../data/BN/FileIndexTable.json")
	json.Unmarshal(raw, &fd)
	hashedFile := FileToMMData(data)
	for i:=0;i<len(fd);i++{
		if reflect.DeepEqual(fd[i].HashedFile, hashedFile){
			index := i
			fd[index].UserID = append(fd[index].UserID, userId)
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
	// db := gormConnect()
  // defer db.Close()
	// for i:=0;i<len(data);i++{
	// 	db.Create(&data[i])
	// }
	// eventsEx := []Log{}
	// db.Find(&eventsEx, "file_id=?", 0)
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

func ParaToJson(data Params, key []PrivKey) {
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

func InputPara() (Params, []PrivKey) {
	var param Params
	var privkeys []PrivKey
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



