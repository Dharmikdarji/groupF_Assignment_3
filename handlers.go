package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new UUID for the ID field
	newItem.ID = uuid.New().String()

	// Save newItem to the database
	stmt, err := db.Prepare("INSERT INTO Items (ID, Name) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newItem.ID, newItem.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve all items from the database
	rows, err := db.Query("SELECT ID, Name FROM Items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to hold the items
	var items []Item

	// Iterate over the rows
	for rows.Next() {
		var item Item

		// Scan the values from the row into the Item struct
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Append the item to the slice
		items = append(items, item)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode items as JSON and respond
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
