package repository

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"TravelPointbackend/internal/utils"
	"fmt"
)

func AuthenticateUser(email, password string) (bool, error) {
	var hashedPassword string

	con := db.OpenConnection()
	err := con.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&hashedPassword)
	if err != nil {
		fmt.Println(err)
	}
	isValid := utils.ComparePasswords(hashedPassword, password)
	db.CloseConnection(con)
	return isValid, err
}

func GetUSerByEmail(email string) (models.User, error) {
	var user models.User
	con := db.OpenConnection()
	err := con.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.LastName, &user.BirthDate, &user.Email, &user.Password, &user.PhoneNumber, &user.IsOwner, &user.CalendarId, &user.AddressId)
	if err != nil {
		fmt.Println(err)
	}
	db.CloseConnection(con)
	return user, err
}
