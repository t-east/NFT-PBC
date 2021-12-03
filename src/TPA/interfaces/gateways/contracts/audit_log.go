package contracts

import ".../domain"

// blog repositoryはSqlHandlerを所持することを宣言
type AuditLogRepository struct {
	// ethネットワークとのつながり，connとかを書く
}

type Params struct {
	Pairing string   `json:"pairing"`
	G       []byte   `json:"g"`
	U       []byte   `json:"u"`
}

func GetPara() (para Params, err error){
	conn, _ := ethhandler.ConnectNetWork()
	reply, err := ethhandler.GetPara(conn)
	if err != nil {
		return
	}
	para := Params{G: reply.G, U: reply.U, Pairing: reply.Pairing}
}

func (repo *AuditLogRepository) AuditChallen() (auditLogs entities.AuditLogs, err error) {
	para, err := GetPara()
	artIds, err := repo.GetArtIds(&bind.CallOpts{})
	if err != nil {
		return
	}
	for i := 0; i < len(artIds); i++ {
		ck, k1, kt2 := pbc.AuditCheallen()
		_, err = repo.SetLogId(auth, uint32(ck), kt1.Bytes(), kt2.Bytes(), []byte(artIds[i]), logId)
		if err != nil {
			return
		}
	}
	var auditLogs entities.AuditLogs
	auditLogs.Logs, err := repo.GetLog()
	if err != nil {
		return
	}
	return
}

func (repo *AuditLogRepository) AuditVerify() (auditLogs entities.AuditLogs, err error) {
	auditLogs := ethhandler.GetLog()
	for adIndex := 0; adIndex < len(auditLogs); p++ {
		hashedFile := ethhandler.GetHashData(auditLogs[adIndex].Log.ArtId)
		auditResult, err := pbc.AuditVerify()
		if err != nil {
			return
		}
		conn, client := ethhandler.ConnectNetWork()
		err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			return
		}
		privKey := os.Getenv("TPA_PRIVATE_KEY")
		auth := ethhandler.AuthUser(client, privKey)
		_, err = conn.SetAuditResult(auth, auditResult, auditLogs[adIndex].LogId )
		if err != nil {
			return
		}
	}
	// 待つ，みたいな処理入れる
	auditLogs = ethhandler.GetLog()
	return
}
