package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"TravelPointbackend/internal/repository"
	"TravelPointbackend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLogin models.UserLogin
	err := c.BindJSON(&userLogin)
	if err != nil {
		fmt.Println(err)
	}
	con := db.OpenConnection()
	var dbUser models.User
	dbUser, err = repository.GetUSerByEmail(userLogin.Email)
	if err != nil {
		fmt.Println(err)
	}
	isValid, err := repository.AuthenticateUser(userLogin.Email, userLogin.Password)
	if err != nil {
		fmt.Println(err)
	}
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	token, err := utils.GenerateJWTToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to genarate token"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
	db.CloseConnection(con)
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func LoginWithGoogle(c *gin.Context) {
	utils.GoogleLogin(c)
}

func GoogleCallback(c *gin.Context) {
	utils.GoogleCallback(c)
}