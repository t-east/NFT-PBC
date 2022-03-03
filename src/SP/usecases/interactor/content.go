package interactor

import (
	entities "pairing_test/src/SP/domains/entities"
	port "pairing_test/src/SP/usecases/port"
)

type ContentHandler struct {
	OutputPort     port.ContentOutputPort
	Repository     port.ContentRepository
	StorageHandler port.ContentStorageHandler
}

// NewUserInputPort はUserInputPortを取得します．
func NewContentInputPort(outputPort port.ContentOutputPort, repository port.ContentRepository, storageHandler port.ContentStorageHandler) port.ContentInputPort {
	return &ContentHandler{
		OutputPort:     outputPort,
		Repository:     repository,
		StorageHandler: storageHandler,
	}
}

func (c *ContentHandler) Create(contentInput *entities.ContentInput) {
	// コンテンツ情報をデータベースに保存
	content, err := c.Repository.Create(contentInput)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	// TODO:ストレージ保存の処理(とりあえずローカルに)
	c.OutputPort.Render(content, 201)
}

func (c *ContentHandler) FindByID(id string) {
	user, err := c.Repository.FindByID(id)
	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}
	c.OutputPort.Render(user, 200)
}
