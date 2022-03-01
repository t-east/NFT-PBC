package controllers

import (
	"net/http"
	"pairing_test/src/user/interfaces/contracts"
	"pairing_test/src/user/usecases/port"

	"gorm.io/gorm"
)

type ContentController struct {
	// -> gateway.NewContentRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	// -> contracts.NewContentContracts
	ContractFactory func() port.ContentContract
	// -> crypt.NewContentCrypt
	CryptFactory func(p contracts.Param) port.ContentCrypt
	// -> presenter.NewContentOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewContentInputPort
	InputFactory func(
		o port.ContentOutputPort,
		u port.ContentRepository,
		cr port.ContentCrypt,
		// co port.ContentContracts,
	) port.ContentInputPort
	Param      contracts.Param
	SQLHandler *gorm.DB
}

func LoadContentController(db *gorm.DB, param contracts.Param) *ContentController {
	return &ContentController{SQLHandler: db, Param: param}
}
