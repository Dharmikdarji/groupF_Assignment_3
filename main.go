package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/items", addItemHandler)
	http.HandleFunc("/items", getItemsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
