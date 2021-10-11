package main

import (
	"pairing_test/src/SP"
	"pairing_test/src/TPA"

	"fmt"
	"os"
	"pairing_test/src/SetUp"

	"pairing_test/src/User"
	"pairing_test/src/tool"
)

// ここの２つの変数を修正して実行する
const (
	// ganacheの起動したときのポートを指定 (8545 か 7545)
	GANACHE_PORT = "8545"
	// 先ほど作成したプログラムから取得した。　CONTRACT_ADDRESSを取得
	CONTRACT_ADDRESS = ""
)

func main() {
	//テスト用にファイルを初期化
	err := os.Remove("../data/BN/FileIndexTable.json")
	err = os.Remove("../data/SP/Storage.json")
	err = os.Remove("../data/BN/FileIndexTable.json")
	err = os.Remove("../data/BN/LogTable.json")
	err = os.Remove("../data/BN/Para.json")
	if err != nil {
		fmt.Println(err)
	}
	//tool.MakeDb()

	//入力データを取得
	dataList := tool.DataListGen()

	//SetUp Phase
	userCount := 3
	params, privKey := SetUp.SetUp(userCount, 160, 512)
	tool.ParaToJson(params, privKey)
	params, privKey = tool.InputPara()

	var challenSs []tool.Chal
	//Upload  Phase
	for i := 0; i < len(dataList); i++ {
		fmt.Printf("\n=====User%dが%sをアップロード======\n", dataList[i].ID, dataList[i].DataName)
		var outsourceData tool.Storage
		outsourceData.MetaData, outsourceData.File  = User.Upload(params, dataList[i], privKey[dataList[i].ID].PrivKey)
		found := User.CheckDep(outsourceData)
		//固有ファイル
		if found == 0 {
			fmt.Printf("\n固有データなので，そのままアウトソース\n")
			outsourceData := User.OutSource(outsourceData)
			SP.SaveOutsourceData(outsourceData, dataList[i].ID)
			//重複ファイル
		} else {
			fmt.Printf("\n重複しているデータなので重複排除処理を行う\n")
			challenS := SP.DedupChallen(params, 10)
			challenSs = append(challenSs,challenS)

			proofS, fitNum := User.DedupProofGen(outsourceData, params, challenS)
			DetupVerifyResult := SP.DedupVerify(outsourceData, params, proofS, challenS, fitNum)
			if DetupVerifyResult == 1 {
				//fmt.Printf("=======DetupVerify===\n       OK\n")
				tool.AddFIT(outsourceData, dataList[i].ID)
			}
		}
	}

	//Auditing Phase
	//challenT := User.AuditChallen(params.Pairing)
	proofT := SP.AuditProofGen(params, challenSs)
	logs, er := TPA.AuditVerify(params, proofT, challenSs)
	if er == 1 {
		fmt.Printf("=====AuditVerify=====\n       OK\n=====================\n")
	} else {
		fmt.Printf("=====AuditVerify=====\n       NG\n=====================\n")
	}
	tool.LTToJson(logs)

	//CheckLog Phase
	for i := 0; i < 3; i++ {
		fmt.Printf("=====User%dのCheckLog=========", i)
		checkResult := User.CheckLog(i, params)
		if checkResult == 1 {
			fmt.Printf("\nUser%dのCheckLogフェーズ成功\n", i)
		} else {
			fmt.Printf("\nUser%dのCheckLogフェーズ失敗\n", i)
		}
	}

}