package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mssql", "server=DESKTOP-7CS9CM5\\SQLEXPRESS;database=StudentAssignment")
	if err != nil {
		log.Fatal(err)
	}
}
