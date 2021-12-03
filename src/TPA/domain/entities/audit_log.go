package entities


type AuditLog struct {
    Result bool `json:"result"`
	C  int    `json:"ck"`
	K1 []byte `json:"k1"`
	K2 []byte `json:"k2"`
    Myu   []byte `json:"myu"`
	Gamma []byte `json:"gamma"`
	ArtId string  `json:"art_id"`
	LogId string  `json:"log_id"`
}

type AuditLogs struct {
    AuditLogs []AuditLog `json:"audit_logs"`
}
