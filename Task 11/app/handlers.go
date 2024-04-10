package app

import (
	"context"
	"csvToMongo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// App struct to hold DB
type App struct {
	mongoDB *mongo.Client
}

// NewApp creates a new App instance with mongoDB client
func NewApp(client *mongo.Client) *App {
	return &App{mongoDB: client}
}

// Function to get the Trains according to the page size and page number
func (app *App) GetAllTrains(w http.ResponseWriter, r *http.Request) {
	pageSize := 10                     // Number of documents per page
	stringPageNum := mux.Vars(r)["id"] // String page Num given by user from URL

	// Convert pageNum to Int
	pageNum, err := strconv.Atoi(stringPageNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Calculate skip value based on page number
	skip := (pageNum - 1) * pageSize

	// Set up options for find operation with paging from skip to skip + pagesize
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))
	// Get data from MongoDB from skip to skip + pagesize
	cursor, err := app.mongoDB.Database("Train").Collection("Train").Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Closing cursor at the end of the function
	defer cursor.Close(context.Background())

	// Iterate the cursor and get all the data in train []models.Train
	var trains []models.Train
	for cursor.Next(context.Background()) {
		var train models.Train
		// Decode JSON data and store in train temporarily
		if err := cursor.Decode(&train); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Append Data in string
		trains = append(trains, train)
	}
	// Marshal the JSON data
	jsonData, err := json.Marshal(trains)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response body
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Function to provide the trains having "findElement" (provide by user in vars) in the train_name field
func (app *App) FindTrains(w http.ResponseWriter, r *http.Request) {
	// Get the Id var and store it in variable findElement
	findElement := mux.Vars(r)["id"]
	// filter query with Regex pattern matching findElement and field table_name
	filter := bson.M{
		"$or": []bson.M{
			{"train_name": bson.M{"$regex": findElement, "$options": "i"}},
			{"starts": bson.M{"$regex": findElement, "$options": "i"}},
			{"ends": bson.M{"$regex": findElement, "$options": "i"}},
		},
	}
	// Set up options for find operation with paging from skip to skip + pagesize
	findOptions := options.Find()
	findOptions.SetLimit(10)
	// Execute the query and fetch the result in cursor
	cursor, err := app.mongoDB.Database("Train").Collection("Train").Find(context.Background(), filter, findOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Closing cursor when function ends
	defer cursor.Close(context.Background())
	// train array to store Data
	var trains []models.Train
	for cursor.Next(context.Background()) {
		// Iterate over cursor
		var train models.Train
		// Decode JSON and add to train temporarily
		if err := cursor.Decode(&train); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// append Data in trains
		trains = append(trains, train)
	}

	// Marshal the data from Trains
	jsonData, err := json.Marshal(trains)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response body and handling Error
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
