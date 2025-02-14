package models

type Address struct {
	ID          string  `json:"id"`
	AddressLine string  `json:"addressLine"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
	PostalCode  string  `json:"postalCode"`
}
