package main

import (
    "net/http"
	// "pairing_test/app/SP"
	// "pairing_test/app/TPA"
	"github.com/gin-gonic/gin"
	// "pairing_test/app/SetUp"

	// "pairing_test/app/User"
	// "pairing_test/app/tool"
)

// func stream() {
// 	//テスト用にファイルを初期化
// 	err := os.Remove("../data/BN/FileIndexTable.json")
// 	err = os.Remove("../data/SP/Storage.json")
// 	err = os.Remove("../data/BN/FileIndexTable.json")
// 	err = os.Remove("../data/BN/LogTable.json")
// 	err = os.Remove("../data/BN/Para.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//入力データを取得
// 	dataList := tool.DataListGen()

// 	//SetUp Phase
// 	// 秘密鍵やParaの設定
// 	userCount := 3
// 	params, privKey := SetUp.SetUp(userCount, 160, 512)
// 	tool.ParaToJson(params, privKey)
// 	params, privKey = tool.InputPara()

// 	//ストレージにファイルを保存，作品ログをブロックチェーンに追加
// 	for i := 0; i < len(dataList); i++ {
// 		fmt.Printf("\n=====User%dが%sをアップロード======\n", dataList[i].ID, dataList[i].DataName)
// 		var outsourceData tool.Storage
// 		outsourceData.MetaData, outsourceData.File  = User.Upload(params, dataList[i], privKey[dataList[i].ID].PrivKey)
// 		found := User.CheckDep(outsourceData)
// 		if found == 0 {
// 			fmt.Printf("\n作品データをアウトソース")
// 			fmt.Printf("\nブロックチェーンにログを追加\n")
// 			outsourceData := User.OutSource(outsourceData)
// 			SP.SaveOutsourceData(outsourceData, dataList[i].ID)
// 		}
// 	}

// 	//Auditing Phase
// 	challenT := User.AuditChallen(params.Pairing)
// 	proofT := SP.AuditProofGen(params, challenT)
// 	logs, er := TPA.AuditVerify(params, proofT, challenT)
// 	if er == 1 {
// 		fmt.Printf("=====AuditVerify=====\n       OK\n=====================\n")
// 	} else {
// 		fmt.Printf("=====AuditVerify=====\n       NG\n=====================\n")
// 	}
// 	//TODO: 監査ログをブロックチェーンに追加
// 	tool.LTToJson(logs)

// 	//CheckLog Phase
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("=====User%dのCheckLog=========", i)
// 		checkResult := User.CheckLog(i, params)
// 		if checkResult == 1 {
// 			fmt.Printf("\nUser%dのCheckLogフェーズ成功\n", i)
// 		} else {
// 			fmt.Printf("\nUser%dのCheckLogフェーズ失敗\n", i)
// 		}
// 	}
// }

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
	  c.String(http.StatusOK, "Hello World!!")
	})
	router.Run(":4000")
  }