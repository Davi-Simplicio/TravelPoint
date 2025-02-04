package models

type Guide struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Rate float64 `json:"rate"`
	AddressId string `json:"addressId"`
}