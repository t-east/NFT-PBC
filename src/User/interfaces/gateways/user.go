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
	if err = ur.SQLHandler.Find(&user, id).Error; err != nil {
		return
	}
	return
}

func (ur *UserRepository) Create(u *entities.User) (err error) {
	err = ur.SQLHandler.Create(u)
	return
}
