package models

type GoogleUser struct {
	ID		  string `json:"id"`
	GoogleID  string `json:"googleId"`
	Name string `json:"name"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	Email     string `json:"email"`
	Verified bool `json:"verified"`
	Picture  string `json:"picture"`
	Provider string `json:"provider"`
	CreatedAt string `json:"createdAt"`
	Password string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	IsOwner bool `json:"isOwner"`
	CalendarID string `json:"calendarId"`
	AddressID string `json:"addressId"`
}