package port

import (
	entities "pairing_test/src/user/domains/entities"
)

type ContentInputPort interface {
	Upload(file *entities.ContentInput)
}

type ContentOutputPort interface {
	Render(*entities.ReceiptFromSP, int)
	RenderError(error)
	UploadSP(*entities.Content) (*entities.ReceiptFromSP, error)
}

type ContentRepository interface {
	Create(user *entities.Content) (*entities.Content, error)
}

type ContentCrypt interface {
	MakeMetaData(contentInput *entities.ContentInput) (*entities.Content, error)
}