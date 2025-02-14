package models

type Chat struct {
	ID       string `json:"id"`
	OwnerId  string `json:"ownerId"`
	ClientId string `json:"clientId"`
	Status   string `json:"status"`
}
