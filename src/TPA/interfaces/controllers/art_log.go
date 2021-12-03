package controller

type ArtLogController struct {
    interactor entities.ArtLogs
}

func NewArtLogController(sqlHandler database.SqlHandler) *ArtLogController {
    return &BlogController {
		interactor: usecase.ArtLogInteractor{
			&database.BlogRepository{
				SqlHandler: sqlHandler
			}
		}
    }
}

func (controller *AuditLogController) Challen(c Context) {
    artIds := entities.ArtIds{}
    c.Bind(&u)
    err := controller.Interactor.Add(u)
    if err != nil {
        c.JSON(500, NewError(err))
        return
    }
    c.JSON(201)
}

func (controller *ArtLogController) AuditChallen(w http.ResponseWriter, r *http.Request) (ret error) {
	blogs, err := controller.interactor.ListBlog()
	ret = response(w, err, map[string]interface{}{"data": blogs})
	return ret
}

func (controller *BlogController) Blog