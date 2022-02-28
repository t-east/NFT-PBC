package interactor

import (
	entities "pairing_test/src/user/domains/entities"
	port "pairing_test/src/user/usecases/port"
)

type ParamHandler struct {
	OutputPort port.ParamOutputPort
	Repository port.ParamRepository
	Crypt  port.ParamCrypt
}

// NewUserInputPort はUserInputPortを取得します．
func NewParamInputPort(outputPort port.ParamOutputPort, repository port.ParamRepository, cryptHandler port.ParamCrypt) port.ParamInputPort {
	return &ParamHandler{
		OutputPort: outputPort,
		Crypt:  cryptHandler,
	}
}

func (ph ParamHandler)GetParam()
