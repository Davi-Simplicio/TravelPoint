package models

type UserEmail struct {
	Email string `json:"email"`
}

type Code struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
