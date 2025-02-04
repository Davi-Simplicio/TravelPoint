package controllers

import (
    "github.com/gin-gonic/gin"
)

import "TravelPointbackend/internal/services"

func Controller(){
	router := gin.Default()
	router.GET("/users", services.GetUsers)
	router.GET("/users/:id", services.GetUserById)
	router.POST("/users", services.PostUers)

	router.GET("/address", services.GetAddress)
	router.GET("/address/:id", services.GetAddressById)
	router.POST("/address", services.PostAddress)

	router.GET("/calendar", services.GetCalendar)
	router.GET("/calendar/:id", services.GetCalendarById)
	router.POST("/calendar", services.PostCalendar)

	router.GET("/date", services.GetDate)
	router.GET("/date/:id", services.GetDateById)
	router.GET("/date/calendar/:calendarId", services.GetDateByCalendarId)
	router.POST("/date", services.PostDate)

	router.Run("localhost:8080")
}



