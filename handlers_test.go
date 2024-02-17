package main

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
)

func setupTest() {
	// Initialize db with a connection to your test database
	// Update the connection string with your test database details
	db, err := sql.Open("sqlserver", "server=DESKTOP-7CS9CM5//SQLEXPRESS;database=Items;integrated security=true")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
}

func TestAddItemHandler(t *testing.T) {
	// Set up test environment
	setupTest()

	// Create a request with JSON payload
	requestBody := []byte(`{"name": "Test Item"}`)
	req, err := http.NewRequest("POST", "/items", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly, passing in the ResponseRecorder and Request
	addItemHandler(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestGetItemsHandler(t *testing.T) {
	// Set up test environment
	setupTest()

	// Create a request
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function directly, passing in the ResponseRecorder and Request
	getItemsHandler(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
