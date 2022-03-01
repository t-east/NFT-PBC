package controllers

import (
	"encoding/json"
	"net/http"
	// "strconv"

	"gorm.io/gorm"

	entities "pairing_test/src/SP/domains/entities"
	port "pairing_test/src/SP/usecases/port"
)

type RequestContent struct {
	MetaData   [][]byte `json:"meta_data"`
	HashedData [][]byte `json:"hashed_data"`
	FileName   string   `json:"name"`
	SplitCount int      `json:"split_count"`
	Owner      string   `json:"owner"`
	ArtId      string   `json:"art_id"`
}

// TODO: ルータやらコントロラーやらハンドラーの関係性が曖昧なので詰める
type ContentController struct {
	// -> presenter.NewUserOutputPort
	OutputFactory func(w http.ResponseWriter) port.ContentOutputPort
	// -> interactor.NewUserInputPort
	InputFactory func(o port.ContentOutputPort, u port.ContentRepository) port.ContentInputPort
	// -> gateway.NewUserRepository
	RepoFactory func(c *gorm.DB) port.ContentRepository
	Conn        *gorm.DB
}

// func NewUsersController(db database.DB) *UsersController {
//     return &UsersController{
//         Interactor: usecase.UserInteractor{
//             DB: &database.DBRepository{ DB: db },
//             User: &database.UserRepository{},
//         },
//     }
// }

// func (cc *ContentController) Get(w http.ResponseWriter, r *http.Request) {

// 	i, _ := ShiftPath(r.URL.Path)
// 	id, err := strconv.Atoi(i)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	repository := cc.RepoFactory(cc.Conn)
// 	outputPort := cc.OutputFactory(w)
// 	inputPort := cc.InputFactory(outputPort, repository)
// 	inputPort.FindByID(id)
// }

func (cc *ContentController) Post(w http.ResponseWriter, r *http.Request) {
	var contentInput entities.ContentInput
	err := json.NewDecoder(r.Body).Decode(&contentInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: ハンドラー，ルータ，メインの変更を反映させる
	outputPort := cc.OutputFactory(w)
	repository := cc.RepoFactory(cc.Conn)
	inputPort := cc.InputFactory(outputPort, repository)
	inputPort.Create(&contentInput)
}
