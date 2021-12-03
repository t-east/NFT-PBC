package controller

type AuditLogController struct {
    interactor entities.AuditLogs
}

func NewAuditLogController(sqlHandler database.SqlHandler) *AuditLogController {
    return &BlogController {
		interactor: usecase.AuditLogInteractor{
			&database.BlogRepository{
				SqlHandler: sqlHandler
			}
		}
    }
}

func (controller *AuditLogController) GetAuditLogs(w http.ResponseWriter, r *http.Request) (ret error) {
	blogs, err := controller.interactor.ListBlog()
	ret = response(w, err, map[string]interface{}{"data": blogs})
	return ret
}

func (controller *BlogController) Blog