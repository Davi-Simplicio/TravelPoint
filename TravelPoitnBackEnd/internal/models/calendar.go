package models

type Calendar struct {
	ID string `json:"id"`
	Date string `json:"date"`
	Availability bool `json:"availability"`
}