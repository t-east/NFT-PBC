package main

import (
	"pairing_test/src/SP"
	"pairing_test/src/TPA"

	//"pairing_test/src/SetUp"
	"fmt"
	"os"

	//"pairing_test/src/TPA"
	"pairing_test/src/User"
	"pairing_test/src/tool"
)

func main() {
	err := os.Remove("../data/BN/FileIndexTable.json")
	err = os.Remove("../data/SP/Storage.json")
	err = os.Remove("../data/fileDataA.json")
	err = os.Remove("../data/fileDataB.json")
	err = os.Remove("../data/fileDataC.json")
	if err != nil {
		fmt.Println(err)
	}
	dataList := tool.DataListGen()
	logTable := []tool.Log{}
	fileDataTable := []tool.FileData{}
	var params tool.Params
	var privKey tool.PrivKeys
	var FileTable []tool.FileIndexTable

	//params, privKey = SetUp.SetUp(4, 160, 512)
	params, privKey = tool.InputPara()
	//tool.ParaToJson(params, privKey)
	for i := 0; i < len(dataList); i++ {
		var fileData tool.FileData
		FileTable = tool.InputFIT()
		found := User.CheckDep(dataList[i], FileTable)
		fileData.DataBlockNum = 10
		fileData.FileName = dataList[i].DataName
		fileData.UserName = dataList[i].Name
		if found == 0 {
			fileData.MetaData, fileData.MData, fileData.MMData, fileData.DataBlockNum, fileData.SplitedData = User.Upload(params, dataList[i], privKey.PrivKeys[i])
			hashedFile, err := tool.HashFile(dataList[i].DataName)
			if err != nil {
				return
			}
			User.OutSource(fileData.SplitedData, fileData.MetaData, dataList[i].DataName, dataList[i])
			var users []string
			users = append(users, dataList[i].Name)
			validFile := tool.FileIndexTable{UserName: users, HashedFile: hashedFile}
			tool.NewFIT(validFile)
			FileTable = append(FileTable, validFile)
		} else {
			fileData.MetaData, fileData.MData, fileData.MMData, fileData.DataBlockNum, fileData.SplitedData = User.Upload(params, dataList[i], privKey.PrivKeys[i])

			challenS := SP.DedupChallen(params.Pairing, fileData.DataBlockNum)

			proofS := User.DedupProofGen(privKey.PrivKeys[i], fileData, params, dataList[i], challenS)

			DetupVerifyResult := SP.DedupVerify(fileData, params, params.PubKeys[i], proofS, challenS)
			if DetupVerifyResult == 1 {
				hashedFile, err := tool.HashFile(dataList[i].DataName)
				if err != nil {
					return
				}
				tool.AddFIT(fileData, hashedFile)
			}
		}
		tool.FileDataToJson(fileData)
		fileDataTable = append(fileDataTable, fileData)
	}

	storage := tool.ReadStorage()
	for i := 0; i < len(storage); i++ {
		challenT := User.AuditChallen(params.Pairing, fileDataTable[i].DataBlockNum)
		proofT := SP.AuditProofGen(fileDataTable[i].DataBlockNum, storage[i], params, challenT)
		log := TPA.AuditVerify(fileDataTable[i].DataBlockNum, params, params.PubKeys[i], storage[i], proofT, challenT)
		if log.Result == 1 {
			fmt.Print("OK")
		}
		logTable = append(logTable, log)
	}
	tool.LTToJson(logTable)

}
