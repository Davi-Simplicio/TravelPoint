package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCalendar(c *gin.Context){
	var con = db.OpenConnection()
	var calendar, err = con.Query("SELECT * FROM calendar")
	if err != nil{
		fmt.Println(err)
	}
	var calendars []models.Calendar
	for calendar.Next(){
		var calendarModel models.Calendar
		err := calendar.Scan(&calendarModel.ID, &calendarModel.Name, &calendarModel.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		calendars = append(calendars, calendarModel)
	}
	c.IndentedJSON(http.StatusOK, calendars)
	db.CloseConnection(con)
}

func GetCalendarById(c *gin.Context){
	id := c.Param("id")
	fmt.Println(id)
	con := db.OpenConnection()
	var calendar models.Calendar
	if err := con.QueryRow("SELECT * FROM calendar WHERE id = $1", id).Scan(&calendar.ID, &calendar.Name, &calendar.Status); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Calendar not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, calendar)
	db.CloseConnection(con)
}

func PostCalendar(c *gin.Context){
	var newCalendar models.Calendar
	sqlStatement := `INSERT INTO calendar (name, status) VALUES ($1, $2) RETURNING id`
	con := db.OpenConnection()
	err := c.BindJSON(&newCalendar)
	if err != nil {
		fmt.Println(err)
	}
	_, err = con.Exec(sqlStatement, newCalendar.Name, newCalendar.Status)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, newCalendar)
	db.CloseConnection(con)
}