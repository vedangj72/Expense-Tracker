package main

import (
	database "ExpenseTacker/Database"
	router "ExpenseTacker/Router"
	// "golang.org/x/net/route"
)

func main() {
	// Initialize the database
	database.InitDatabase()

	// Setup the router
	router := router.SetupRouter()

	router.Run(":8080")
}
