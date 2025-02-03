package db

import (
    "database/sql"
    "fmt"
  
    _ "github.com/lib/pq"
  )

  const (
		host = "localhost"
	port = 5433
	user = "postgres"
	password = "Lau27439635vi@"
	dbname = "travel_point"
)

func OpenConnection() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db

}

func CloseConnection(db *sql.DB) {
	db.Close()
	db.Ping()
	fmt.Println("Connection closed!")
}