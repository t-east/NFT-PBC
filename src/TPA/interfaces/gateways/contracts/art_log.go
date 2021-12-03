package contracts

import ".../domain"

// blog repositoryはSqlHandlerを所持することを宣言
type ArtLogRepository struct {
	// ethネットワークとのつながり，connとかを書く
}

func (repo *ArtLogRepository) Count() (artLogs entities.ArtLogs, err error) {
	artIdBytes, err := repo.GetArtIds(&bind.CallOpts{})
	if err != nil {
		return
	}
	var artLogs entities.ArtLogs
	artLogs.Ids = artIdBytes
}
