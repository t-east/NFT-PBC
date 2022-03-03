package gateways

import (
	"pairing_test/src/user/domains/entities"
	"pairing_test/src/user/usecases/port"

	"gorm.io/gorm"
)

type SQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.User, error)
	First(interface{}, ...interface{}) (*entities.User, error)
	Create(interface{}) error
	Save(interface{}) error
	Delete(interface{}) *entities.User
	Where(interface{}, ...interface{}) *entities.User
}

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *port.UserRepository {
	return &UserRepository{
		Conn: conn,
	}
}

func (ur *UserRepository) FindByID(id int) (user *entities.User, err error) {
	userInDB, err := ur.Handler.Find(&user, id)
	if err != nil {
		return nil, err
	}
	return userInDB, nil
}

func (ur *UserRepository) Create(u *entities.User) (user *entities.User, err error) {
	err = ur.Handler.Create(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) Update(u *entities.User) (user *entities.User, err error) {
	err = ur.Handler.Save(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}
