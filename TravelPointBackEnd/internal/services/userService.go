package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context){
	var con = db.OpenConnection()
	var user, err = con.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	var users []models.User
	for user.Next(){
		var userModel models.User
		err := user.Scan(&userModel.ID, &userModel.Name, &userModel.LastName, &userModel.BirthDate, &userModel.Email, &userModel.Password, &userModel.PhoneNumber, &userModel.IsOwner, &userModel.CalendarId, &userModel.AddressId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, userModel)
	}
	c.IndentedJSON(http.StatusOK, users)
	db.CloseConnection(con)
}

func GetUserById(c *gin.Context){
	id := c.Param("id")

	for _, a:= range db.Users{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func PostUers(c *gin.Context){
	var newUser models.User
	sqlStatement := `INSERT INTO users (name, lastName, birthDate, email, password, phoneNumber, isOwner, calendarId, addressId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	con := db.OpenConnection()
	err := c.BindJSON(&newUser)
	print(err)
	if err != nil {
		fmt.Println(err)
	}
	_, err = con.Exec(sqlStatement, newUser.Name, newUser.LastName, newUser.BirthDate, newUser.Email, newUser.Password, newUser.PhoneNumber, newUser.IsOwner, newUser.CalendarId, newUser.AddressId)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, newUser)
	db.CloseConnection(con)
}