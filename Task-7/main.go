package main

import (
	"log"
	"net/http"

	"github.com/Tridipchavda/middleware/api"
	"github.com/Tridipchavda/middleware/dbconn"
)

// Main function to Start DB and Listen for HTTP requests
func main() {
	// Create the Database connection and store the DB connection to execute Query
	db, err := dbconn.InitDB()
	if err != nil {
		log.Println("Error in Database Initialization", err)
		return
	}
	// Make new api instanace
	myapi := api.NewAPI(db)

	// Handling Routes
	http.HandleFunc("/generate", myapi.Sign)
	http.Handle("/login", myapi.AuthenticateToken(myapi.LogIn))

	// Starting the server at PORT 8080
	log.Println("Server start at 8080")
	http.ListenAndServe(":8080", nil)
}
