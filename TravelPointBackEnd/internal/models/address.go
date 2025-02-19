package models

type Address struct {
	ID           string  `json:"id"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Cep          string  `json:"cep"`
	Neighborhood string  `json:"neighborhood"`
	Street       string  `json:"street"`
	Number       string  `json:"number"`
}
