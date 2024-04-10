package main

import (
	"context"
	"csvToMongo/app"
	"csvToMongo/utils"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Variable to store the filename for data Entry
	var fileName string
	// Store the mongoClient in variable
	mongodb := utils.ConnMongoDB()
	// Closing the connection at the end
	defer mongodb.Disconnect(context.Background())
	// Store the variable from flag in cmd to fileName
	flag.StringVar(&fileName, "insertDB", "-", "CSV File to write in the database")
	flag.Parse()
	// check if file is entered by user
	if fileName != "-" {
		// Insert Data from csv to Mongo database if file exist and handle error
		log.Println("Inserting Data in database")
		err := utils.InsertDataFromCSV(fileName, mongodb)
		if err != nil {
			log.Println(err)
		}
	} else {
		// Create instance of App to pass mongoDB in it
		app := app.NewApp(mongodb)
		// Initalize Mux Router
		r := mux.NewRouter()

		// Handle Routes using mux for get paging Data and filtered Data
		r.HandleFunc("/find/{id}", app.FindTrains).Methods("GET")
		r.HandleFunc("/{id}", app.GetAllTrains).Methods("GET")

		// Starting the server at 8080
		log.Println("Listening Server on Port 8080")
		http.ListenAndServe(":8080", r)
	}
}
