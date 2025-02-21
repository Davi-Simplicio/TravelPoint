package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"TravelPointbackend/internal/repository"
	"TravelPointbackend/internal/utils"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var codeStore = make(map[string]string)
var codeMutex = sync.Mutex{}

func GetUsers(c *gin.Context) {
	var con = db.OpenConnection()
	var user, err = con.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	var users []models.User
	for user.Next() {
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

func GetUserById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range db.Users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func PostUers(c *gin.Context) {
	var newUser models.User
	sqlStatement := `INSERT INTO users (name, lastName, birthDate, email, password, phoneNumber, isOwner, calendarId, addressId) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	con := db.OpenConnection()
	err := c.BindJSON(&newUser)
	print(err)
	if err != nil {
		fmt.Println(err)
	}
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), cost)
	if err != nil {
		fmt.Println(err)
	}
	newUser.Password = string(hash)

	calendarName := newUser.Name + " " + "Calendar"

	userCalendar := models.Calendar{
		Name:   calendarName,
		Status: true,
	}
	userCalendar, err = PostCalendar(userCalendar)
	if err != nil {
		fmt.Println(err)
	}
	newUser.CalendarId, err = strconv.Atoi(userCalendar.ID)

	if err != nil {
		fmt.Println(err)
	}

	userErr := con.QueryRow(sqlStatement, newUser.Name, newUser.LastName, newUser.BirthDate, newUser.Email, newUser.Password, newUser.PhoneNumber, newUser.IsOwner, newUser.CalendarId, newUser.AddressId).Scan(&newUser.ID)
	if userErr != nil {
		fmt.Println("é aqui")
		fmt.Println(userErr)
	}
	c.IndentedJSON(http.StatusCreated, newUser)
	db.CloseConnection(con)
}

func RetrieveToken(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	tokenDecoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenDecoded})
}

func ChangePassword(c *gin.Context) {
	var userLogin models.UserLogin
	err := c.BindJSON(&userLogin)
	if err != nil {
		fmt.Println(err)
	}
	con := db.OpenConnection()

	isValid, err := repository.AuthenticateUser(userLogin.Email, userLogin.Password)

	if err != nil {
		fmt.Println(err)
	}
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	var user models.User
	user, err = repository.GetUSerByEmail(userLogin.Email)
	if err != nil {
		fmt.Println(err)
	}

	var cost = bcrypt.DefaultCost

	hash, err := bcrypt.GenerateFromPassword([]byte(userLogin.Password), cost)

	if err != nil {
		fmt.Println(err)
	}

	user.Password = string(hash)

	_, err = con.Query("UPDATE users SET password = $1 WHERE email = $2", user.Password, user.Email)

	if err != nil {
		fmt.Println(err)
	}

	db.CloseConnection(con)
}

func SendCodeByEmail(c *gin.Context) {
	var userEmail models.UserEmail
	err := c.BindJSON(&userEmail)
	if err != nil {
		fmt.Println(err)
	}
	codeMutex.Lock()
	codeStore[userEmail.Email], err = utils.SendCode(userEmail.Email)
	if err != nil {
		fmt.Println(err)
	}
	codeMutex.Unlock()
}

func AuthenticateCode(c *gin.Context) {
	var userCode models.Code
	err := c.BindJSON(&userCode)
	if err != nil {
		fmt.Println(err)
	}
	print(codeStore[userCode.Email])
	if utils.AuthenticateCode(codeStore[userCode.Email], userCode.Code) {
		c.JSON(http.StatusOK, gin.H{"message": "Code is valid"})
		codeMutex.Lock()
		delete(codeStore, userCode.Email)
		codeMutex.Unlock()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Code is invalid"})
	}
}
