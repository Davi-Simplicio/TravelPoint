package main

import "fmt"
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	LastName string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	Email string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	IsOwner bool `json:"isOwner"`
	CalendarId string `json:"calendarId"`
	AddressId string `json:"addressId"`
}