package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	BirthDate   string `json:"birthDate"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	IsOwner     bool   `json:"isOwner"`
	CalendarId  int    `json:"calendarId"`
	AddressId   int    `json:"addressId"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
