package models

type Date struct {
	ID         string `json:"id"`
	Date       string `json:"date"`
	CalendarId string `json:"calendarId"`
	HasEvent   bool   `json:"hasEvent"`
}
