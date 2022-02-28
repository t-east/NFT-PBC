package interactor

import (
	entities "pairing_test/src/user/domains/entities"
	port "pairing_test/src/user/usecases/port"
)

type UserHandler struct {
	InputPort port.UserOutputPort
	OutputPort port.UserOutputPort
	Repository port.UserRepository
	Crypt port.UserCrypt
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, repository port.UserRepository) port.UserInputPort {
	return &UserHandler{
		OutputPort: outputPort,
		Repository:  repository,
	}
}

func (uc *UserHandler) Create(ui port.UserCreate) {
	// TODO: ユーザ情報を確認する -> 登録済みなのに新しく鍵を生成したら色々狂う

	// 鍵作成
	key, err := uc.Crypt.KeyGen()
	if err != nil {
		uc.OutputPort.RenderError(err)
		return
	}
	// entitiesの型に変換
	// TODO: バリデーションを付ける
	user := &entities.User{
		Address: ui.Address,
		PubKey: key.PubKey,
		PrivKey: key.PrivKey,
	}

	// user情報を保存する
	user, err = uc.Repository.Create(user)
	if err != nil {
		uc.OutputPort.RenderError(err)
		return
	}

	uc.OutputPort.Render(user, 201)
}