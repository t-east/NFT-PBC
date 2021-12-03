package server

/*
 * Input Port
 *  └─ Interactor で実装、Controller で使用される
 */
type ArtIdsInputPort interface {
  	ArtIds() (*ArtIdsResponse, port.Error)
}

type ArtIds struct {
  ArtIds entities.ArtIdsc
}

/*
 * Output Port
 *  └─ Presenter で実装、Interactor で使用される
 */
type ArtIdsOutputPort interface {
	ArtIds(entities.ArtIds) (*ArtIdsResponse, port.Error)
}

type ArtIdsResponse struct {
  DataSource *entities.ArtIds
}