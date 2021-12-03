package repository

type ArtIdsRepository interface {
  Count() (entities.ArtIds, port.Error)
}