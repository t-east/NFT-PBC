package port

import (
	entities "pairing_test/src/user/domains/entities"
)

type ParamInputPort interface {
	GetParam(file *entities.Content)
}

type ParamOutputPort interface {
	Render(*entities.Param, int)
	RenderError(error)
}

type ParamBC interface {
	GetParam(*entities.User) (*entities.Param, error)
}