package models

type Owner struct {
	ID string `json:"id"`
	UserId string `json:"userId"`
	Earnings float64 `json:"earnings"`
}