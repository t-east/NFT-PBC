package entities

type User struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	PubKey  string `json:"owner"`
	PrivKey string `json:"art_id"`
}
