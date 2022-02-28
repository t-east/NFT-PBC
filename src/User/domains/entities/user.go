package entities

type User struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}

type Key struct {
	PubKey  string `json:"pub_key"`
	PrivKey string `json:"priv_key"`
}
