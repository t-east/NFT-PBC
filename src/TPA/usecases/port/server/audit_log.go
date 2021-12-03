package server

/*
 * Input Port
 *  └─ Interactor で実装、Controller で使用される
 */
type AuditLogsInputPort interface {
  AuditLogs() (*AuditLogsResponse, port.Error)
}

type AuditLogs struct {
  AuditLogs entities.AuditLogs
}

/*
 * Output Port
 *  └─ Presenter で実装、Interactor で使用される
 */
type AuditLogsOutputPort interface {
	AuditLogs(entities.AuditLogs) (*AuditLogsResponse, port.Error)
}

type AuditLogsResponse struct {
  DataSource *entities.AuditLogs
}