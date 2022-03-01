package controllers

import (
	"net/http"
	"pairing_test/src/user/interfaces/contracts"
	"pairing_test/src/user/usecases/port"
	"gorm.io/gorm"
)

type UserController struct {
	// -> gateway.NewUserRepository
	RepoFactory func(c *gorm.DB) port.UserRepository
	// -> crypt.NewUserCrypt
	CryptFactory func(p contracts.Param) port.UserCrypt
	// -> presenter.NewUserOutputPort
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> interactor.NewUserInputPort
	InputFactory func(
		o port.UserOutputPort,
		u port.UserRepository,
		cr port.UserCrypt,
		// co port.UserContracts,
	) port.UserInputPort
	Param contracts.Param
	Conn        *gorm.DB
}

func LoadUserController(db *gorm.DB, param contracts.Param) *UserController {
	return &UserController{Conn: db, Param: param}
}
