package services

import (
	"TravelPointbackend/internal/db"
	"TravelPointbackend/internal/models"
	"TravelPointbackend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddress(c *gin.Context) {
	var con = db.OpenConnection()
	var address, err = con.Query("SELECT * FROM address")
	if err != nil {
		panic(err)
	}

	var addresses []models.Address

	for address.Next() {
		var addressModel models.Address
		err := address.Scan(&addressModel.ID,  &addressModel.Longitude, &addressModel.Latitude, &addressModel.City, &addressModel.State, &addressModel.Country, &addressModel.Cep, &addressModel.Neighborhood, &addressModel.Street, &addressModel.Number)
		if err != nil {
			fmt.Println(err)
			continue
		}
		addresses = append(addresses, addressModel)
	}
	c.IndentedJSON(http.StatusOK, addresses)
	db.CloseConnection(con)
}

func GetAddressById(c *gin.Context) {
	var con = db.OpenConnection()
	id := c.Param("id")
	var address models.Address
	if err := con.QueryRow("SELECT * FROM address WHERE id = $1", id).Scan(&address.ID, &address.Longitude, &address.Latitude, &address.City, &address.State, &address.Country, &address.Cep, &address.Neighborhood, &address.Street, &address.Number ); err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, address)
	db.CloseConnection(con)
}

func PostAddress(c *gin.Context) {
	var newAddress models.Address
	sqlStatement := `INSERT INTO address (longitude, latitude, city, state, country, cep, neighborhood, street, number) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	con := db.OpenConnection()
	err := c.BindJSON(&newAddress)
	if err != nil {
		fmt.Println(err)
	}
	AddressErr := con.QueryRow(sqlStatement,  newAddress.Longitude, newAddress.Latitude, newAddress.City, newAddress.State, newAddress.Country, newAddress.Cep, newAddress.Neighborhood, newAddress.Street, newAddress.Number).Scan(&newAddress.ID)

	if AddressErr != nil {
		fmt.Println(AddressErr)
	}
	fmt.Println(newAddress.ID)
	c.IndentedJSON(http.StatusCreated, newAddress)
	db.CloseConnection(con)
}

func GetAddressByCep(c *gin.Context) {
	cep := c.Param("cep")
	address, err := utils.GetInfoByCep(cep)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, address)
}
