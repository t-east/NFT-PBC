package gateway

import (
	"pairing_test/src/user/domains/entities"
	"pairing_test/src/user/drivers"
	"pairing_test/src/user/usecases/port"
	"pairing_test/src/user/interfaces/contracts"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func(c *gorm.DB) port.ContentContracts
	// -> crypt.NewContentCrypt
	CryptFactory func(c *gorm.DB) port.ContentCrypt
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		cr port.ContentCrypt,
		co port.ContentContracts
	) port.UserInputPort
	Param contracts.Param
	Conn        *gorm.DB
}

func LoadContentController(db *gorm.DB, param *contracts.Param) *ContentController {
	cc := &ContentController{Conn: db, Param: param}
	return cc
}
