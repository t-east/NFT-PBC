package port

import (
	entities "pairing_test/src/user/domains/entities"
)

type UserCreate struct {
	Name string
	Address string
}

type UserInputPort interface {
	Create(UserCreate)
}

type UserOutputPort interface {
	Render(*entities.User, int)
	RenderError(error)
}

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
}

type UserCrypt interface {
	KeyGen() (*entities.Key, error)
}