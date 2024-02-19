package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // Import the MSSQL driver package
)

// Item represents a single item in the collection
type Item struct {
	ID   string `json:"id"`   // ID of the item
	Name string `json:"name"` // Name of the item
}

var db *sql.DB // Global variable to hold the database connection

// initDB initializes the database connection
func initDB() {
	var err error
	// Open a connection to the SQL Server database
	db, err = sql.Open("mssql", "server=DESKTOP-7CS9CM5\\SQLEXPRESS;database=StudentAssignment")
	if err != nil {
		log.Fatal(err) // If there's an error opening the connection, log and exit
	}
}
