package port

import (
	entities "pairing_test/src/user/domains/entities"
)

type UserInputPort interface {
	Create(*entities.User)
	KeyGen(int)
	Get(int)
}

type UserOutputPort interface {
	Render(*entities.User, int)
	RenderError(error)
}

type UserRepository interface {
	Create(*entities.User) (*entities.User, error)
	Update( *entities.User) (*entities.User, error)
	FindById(int) (*entities.User, error)
}

type UserCrypt interface {
	KeyGen() (*entities.Key, error)
}