# Simple Item Management API

This is a simple RESTful API written in Go that manages a collection of items. Each item has an ID and a Name. The API provides functionality to add new items and retrieve the list of all items.

## API Endpoints

### POST /items
Adds a new item to the collection.
Request body should contain the name of the item.
Generates an ID for the new item.
Returns the created item with its ID.

### GET /items
Retrieves a list of all items.



### Database Configuration
The API uses a SQL Server database for storing items. To configure the database connection, modify the initDB function in main.go with the appropriate connection string.

### Usage
Make sure you have Go installed on your machine.
Clone this repository to your local machine.
Navigate to the project directory.
Run the following commands to build and run the API:
go build
The API will start running on http://localhost:8080.


### Dependencies
github.com/denisenkom/go-mssqldb - MSSQL driver for Go.