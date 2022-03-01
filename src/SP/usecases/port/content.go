package port

import (
	entities "pairing_test/src/SP/domains/entities"
)

type ContentInputPort interface {
	FindByID(id int)
	Create(file *entities.ContentInput)
}

type ContentOutputPort interface {
	Render(*entities.Content, int)
	RenderError(error)
	// ファイルをそのまま返す
	// RenderContent()
}

// userのCRUDに対するDB用のポート
type ContentRepository interface {
	Create(user *entities.ContentInput) (*entities.Content, error)
	FindByID(id int) (*entities.Content, error)
}

// ストレージを利用するport
type ContentStorageHandler interface {
	Create(user *entities.ContentInStorage) error
	FindByID(id string) (*entities.ContentInStorage, error)
}
