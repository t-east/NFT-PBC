package interactor

import (
	entities "pairing_test/src/user/domains/entities"
	port "pairing_test/src/user/usecases/port"
)

type ContentHandler struct {
	OutputPort port.ContentOutputPort
	Repository port.ContentRepository
	Crypt  port.ContentCrypt
}

// NewUserInputPort はUserInputPortを取得します．
func NewContentInputPort(outputPort port.ContentOutputPort, repository port.ContentRepository, cryptHandler port.ContentCrypt) port.ContentInputPort {
	return &ContentHandler{
		OutputPort: outputPort,
		Repository: repository,
		Crypt:  cryptHandler,
	}
}

func (c *ContentHandler) Upload(contentInout *entities.ContentInput) {
	// メタデータ作成
	content, err := c.Crypt.MakeMetaData(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	// ArtIDをデータベースに保存
	content, err = c.Repository.Create(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	// SPにアップロードする
	receipt, error := c.OutputPort.UploadSP(content)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	c.OutputPort.Render(receipt)
}
