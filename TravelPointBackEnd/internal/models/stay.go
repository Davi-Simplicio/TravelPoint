package models

type Stay struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	RentalPeriod string  `json:"rentalPeriod"`
	Description  string  `json:"description"`
	Rate         float64 `json:"rate"`
	MaxGuests    int     `json:"maxGuests"`
	OwnerId      string  `json:"ownerId"`
	CalendarId   string  `json:"calendarId"`
	AddressId    string  `json:"addressId"`
}
