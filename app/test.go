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

func main() {
	//ストレージ初期化
	err := os.Remove("../data/SP/Storage.json")
	err := os.Remove("../data/BN/Para.json")
	if err != nil {
		fmt.Println(err)
	}

	//テスト用入力データを取得
	dataList := tool.DataListGen()

	//SetUp Phase
	// 秘密鍵やParaの設定
	userCount := 3
	params, privKey := SetUp.SetUp(userCount, 160, 512)
	// tool.ParaToJson(params, privKey)
	// paramsとprivKeyから生成したpubKeyをブロックチェーンに登録
	SetUp.addPara(params, privKey)
	params, privKey = tool.InputPara()

	//ストレージにファイルを保存，作品ログをブロックチェーンに追加
	for i := 0; i < len(dataList); i++ {
		fmt.Printf("\n=====User%dが%sをアップロード======\n", dataList[i].ID, dataList[i].DataName)
		var outsourceData tool.Storage
		outsourceData.MetaData, outsourceData.File  = User.Upload(params, dataList[i], privKey[dataList[i].ID].PrivKey)
		fmt.Printf("\n作品データをアウトソース")
		fmt.Printf("\nブロックチェーンにログを追加\n")
		// outsourceDataはuserがSPに送る値(ファイルとメタデータとuserIDが必要)
		outsourceData := User.OutSource(outsourceData)
		// outsourceDataをストレージに保存してブロックチェーンに情報を書き込む
		// ユーザに返す作品ログのartIdを出力
		artId := SP.SaveOutsourceData(outsourceData, dataList[i].ID)
		fmt.Println(artId)
	}

	//Auditing Phase
	// SPに送るチャレンジノンスを作成して出力している
	challenT := User.AuditChallen(params.Pairing)
	// チャレンジノンスを用いて証明データ作成
	proofT := SP.AuditProofGen(params, challenT)
	// 監査ログに書き込む値をlogsとして出力
	logs, er := TPA.AuditVerify(params, proofT, challenT)
	if er == 1 {
		fmt.Printf("=====AuditVerify=====\n       OK\n=====================\n")
	} else {
		fmt.Printf("=====AuditVerify=====\n       NG\n=====================\n")
	}
	//logsをブロックチェーンに追加
	// TPA.addLog(logs)

	//CheckLog Phase
	// ユーザは作品ログのidを入力するだけで確認ができる
	fmt.Printf("=====User%dのCheckLog=========", i)
	checkResult := User.CheckLog(artId,params)
	if checkResult == 1 {
		fmt.Printf("\nUser%dのCheckLogフェーズ成功\n", i)
	} else {
		fmt.Printf("\nUser%dのCheckLogフェーズ失敗\n", i)
	}
}