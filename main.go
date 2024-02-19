package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize the database connection
	initDB()

	// Register HTTP request handlers for "/items" endpoint
	http.HandleFunc("/items", addItemHandler)  // Handler for adding new items
	http.HandleFunc("/items", getItemsHandler) // Handler for retrieving all items

	// Start the HTTP server and listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
