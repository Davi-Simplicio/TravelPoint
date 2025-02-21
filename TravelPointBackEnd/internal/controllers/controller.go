package controllers

import (
	"TravelPointbackend/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Controller() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin"},
		AllowCredentials: true,
	}))

	// User Service
	router.GET("/users", services.GetUsers)
	router.GET("/users/:id", services.GetUserById)
	router.POST("/users", services.PostUers)
	router.GET("/token", services.RetrieveToken)

	// Login Service
	router.POST("/login", services.Login)
	router.GET("/logout", services.Logout)
	router.GET("/login/google", services.LoginWithGoogle)
	router.GET("/google/callback", services.GoogleCallback)
	
	// Address Service
	router.GET("/address", services.GetAddress)
	router.GET("/address/:id", services.GetAddressById)
	router.GET("/address/cep/:cep", services.GetAddressByCep)
	router.POST("/address", services.PostAddress)

	// Calendar Service
	router.GET("/calendar", services.GetCalendar)
	router.GET("/calendar/:id", services.GetCalendarById)

	// Date Service
	router.GET("/date", services.GetDate)
	router.GET("/date/:id", services.GetDateById)
	router.GET("/date/calendar/:calendarId", services.GetDateByCalendarId)
	router.POST("/date", services.PostDate)

	// Email Verification Code
	router.POST("/sendCode", services.SendCodeByEmail)
	router.POST("/verifyCode", services.AuthenticateCode)

	router.Run("localhost:8080")
}
