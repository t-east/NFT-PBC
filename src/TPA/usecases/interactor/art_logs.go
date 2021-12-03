package interactor

import (
  ".../sample-app/config" 
  ".../sample-app/usecase/port"
  ".../sample-app/usecase/port/server"
  ".../sample-app/usecase/port/repository"
)

type ArtIdsInteractor struct {
  Config *config.Config
  OutputPort server.ArtIdsOutputPort
  Repository repository.ArtIdsRepository
}

func NewArtIdsInteractor(
  config *config.Config,
  outputPort server.DataSourceOutputPort,
  dataSourceRepository repository.DataSourceRepository,
) *ArtIdsInteractor {
  return &ArtIdsInteractor{
    Container: newUsecaseContainer(config),
    OutputPort: outputPort,
    DataSourceRepository: dataSourceRepository,
  }
}

// Input Port の実装
func (i *ArtIdsInteractor) Count() (*server.DownloadDataSourcesResponse, port.Error) {
  res, err := i.Repository.Count()
  if err != nil {
    return nil, err
  }
  // Output Port の使用
  return i.OutputPort.Count(res)
}
