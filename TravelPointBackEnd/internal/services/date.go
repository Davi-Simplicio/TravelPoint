package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDate(c *gin.Context){
	var con = db.OpenConnection()
	var date, err = con.Query("SELECT * FROM date")
	if err != nil {
		fmt.Println(err)
	}
	var dates []models.Date
	for date.Next(){
		var dateModel models.Date;
		err := date.Scan(&dateModel.ID, &dateModel.Date, &dateModel.CalendarId, &dateModel.HasEvent)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dates = append(dates, dateModel)
	}
	c.IndentedJSON(http.StatusOK, dates)
	db.CloseConnection(con)
}

func GetDateById(c *gin.Context){
	id := c.Param("id")
	var con = db.OpenConnection()
	var date models.Date
	if err := con.QueryRow("SELECT * FROM date WHERE id = $1", id).Scan(&date.ID, &date.Date, &date.CalendarId, &date.HasEvent); err != nil {
		fmt.Println(err)
	} 
	c.IndentedJSON(http.StatusOK, date)
	db.CloseConnection(con)
}

func PostDate(c *gin.Context){
	var newDate models.Date
	sqlStatement := `INSERT INTO date (date, calendarId, hasEvent) VALUES ($1, $2, $3) RETURNING id`
	var con = db.OpenConnection()
	err := c.BindJSON(&newDate)
	if err != nil {
		fmt.Println(err)
	}
	dateErr := con.QueryRow(sqlStatement, newDate.Date, newDate.CalendarId, newDate.HasEvent).Scan(&newDate.ID)
	if dateErr != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, newDate)
	db.CloseConnection(con)
}

func GetDateByCalendarId(c *gin.Context){
	id := c.Param("calendarId")
	var con = db.OpenConnection()
	var date models.Date
	if err := con.QueryRow("SELECT * FROM date WHERE calendarId = $1", id).Scan(&date.ID, &date.Date, &date.CalendarId, &date.HasEvent); err != nil {
		fmt.Println(err)
	} 
	c.IndentedJSON(http.StatusOK, date)
	db.CloseConnection(con)
}