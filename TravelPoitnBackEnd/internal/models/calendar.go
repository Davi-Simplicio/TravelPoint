package models

type Calendar struct {
	ID string `json:"id"`
	EntityType string `json:"entityType"`
	EntityId string `json:"entityId"`
	Date string `json:"date"`
	Availability string `json:"availability"`
}