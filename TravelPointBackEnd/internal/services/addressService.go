package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddress(c *gin.Context){
	var con = db.OpenConnection()
	var address, err = con.Query("SELECT * FROM address")
	if err != nil{
		panic(err)
	}

	var addresses []models.Address

	for address.Next(){
		var addressModel models.Address
		err := address.Scan(&addressModel.ID,&addressModel.AddressLine, &addressModel.Latitude, &addressModel.Longitude, &addressModel.City, &addressModel.State, &addressModel.Country, &addressModel.PostalCode)
		if err != nil {
			fmt.Println(err)
			continue
		}
		addresses = append(addresses, addressModel)
	}
	c.IndentedJSON(http.StatusOK, addresses)
	db.CloseConnection(con)
}

func GetAddressById(c *gin.Context){
	var con = db.OpenConnection()
	id := c.Param("id")
	var address models.Address
	if err := con.QueryRow("SELECT * FROM address WHERE id = $1",id).Scan(&address.ID, &address.AddressLine, &address.Latitude, &address.Longitude, &address.City, &address.State, &address.Country, &address.PostalCode); 
	err != nil{
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, address)
	db.CloseConnection(con)
}

func PostAddress(c * gin.Context){
	var newAddress models.Address
	sqlStatement := `INSERT INTO address (addressLine, latitude, longitude, city, state, country, postalCode) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	con := db.OpenConnection()
	err := c.BindJSON(&newAddress)
	if err != nil {
		fmt.Println(err)
	}	
	_, err = con.Exec(sqlStatement, newAddress.AddressLine, newAddress.Latitude, newAddress.Longitude, newAddress.City, newAddress.State, newAddress.Country, newAddress.PostalCode)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusCreated, newAddress)
	db.CloseConnection(con)
}