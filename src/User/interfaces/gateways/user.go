package gateways

import (
	"pairing_test/src/user/domains/entities"
	"pairing_test/src/user/drivers/rdb"
	"pairing_test/src/user/usecases/port"

	"gorm.io/gorm"
)

type UserRepository struct {
	SQLHandler
}

func NewUserRepository(conn *gorm.DB) port.UserRepository {
	sqlHandler := &rdb.SQLHandler{Conn: conn}
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

func (ur *UserRepository) FindByID(id int) (user *entities.User, err error) {
	userInDB, err := ur.SQLHandler.Find(&user, id)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (ur *UserRepository) Create(u *entities.UserCreate) (*entities.User, error) {
	user, err = ur.SQLHandler.Create(u)
	return user, err
}
