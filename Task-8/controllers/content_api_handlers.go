package controllers

import (
	"context"
	"errors"
	"strconv"

	"github.com/Tridipchavda/FiberWithMongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	{
	  "id":123,
	  "name":"Oppenheimer",
	  "type":"movie",
	  "cast": {
	    "directors":["Christopher Nolan"],
	    "actors":["Cillian Murphy","RDJ"],
	    "producers":["Christoper Nolan","Emma Thomas"],
	    "totalCrew": 327
	  }
	  "genre":"sci-fi",
	  "release_date":"21-10-2023",
	  "category": "Hollywood",
	  "trailer" : "https://www.youtube.com/watch?v=uYPbbksJxIg"
	}
*/

// Controller struct with DB connection
type Controller struct {
	DB *mongo.Client

}

// Function to Initialize the controller struct
func NewController(db *mongo.Client) *Controller {
	return &Controller{DB: db}
}

// Function to get all the Content
func (con *Controller) GetAllContentData() ([]models.Content, error) {

	// Query MongoDB for Read all the Content
	var contents []models.Content
	cur, err := con.DB.Database("Golang").Collection("Content").Find(context.Background(),bson.M{})

	// Handling Error
	if err != nil {
		return nil, err
	}

	// Setting Up the Data from cursor to the Content array 
	for cur.Next(context.Background()) {
		var content models.Content
		cur.Decode(&content)
		contents = append(contents, content)
	}

	// Send Data to Route Handlers
	return contents, nil
}

// Function to get one of the Content
func (con *Controller) GetOneContentData(id string) (models.Content, error) {
	// Convert id to string from string format
	i, _ := strconv.Atoi(id)

	// Fetch the document by Id and store in content
	var content models.Content
	con.DB.Database("Golang").Collection("Content").FindOne(context.Background(), bson.M{"id": i}).Decode(&content)

	// return data from MongoDB
	return content, nil
}

// Function to Insert Data in MongoDB
func (con *Controller) InsertContent(content *models.Content) error {
	// Insert Document by fetching the JSON in content struct in arguments
	_, err := con.DB.Database("Golang").Collection("Content").InsertOne(context.Background(), content)

	// Handling the error
	if err != nil {
		return err
	}

	return nil
}

// Function to update the Document in MongoDB
func (con *Controller) UpdateContent(content models.Content, id string) error {
	// Convert the id to int from string
	i, _ := strconv.Atoi(id)

	// Update the Record by fetching the id
	res, err := con.DB.Database("Golang").Collection("Content").UpdateOne(context.Background(), bson.M{"id": i}, bson.M{"$set": content})

	// Handling the error from Database
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("ERROR: No such ID Found ")
	}
	return nil
}

// Function to delete the Document in MongoDB by Id
func (con *Controller) DeleteContent(id string) error {
	// Convert the id to int from string
	i, _ := strconv.Atoi(id)

	// Delete the Document from MongoDb by Id
	res, err := con.DB.Database("Golang").Collection("Content").DeleteMany(context.Background(), bson.M{"id": i}, nil)

	// Handling Error from Database
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("ERROR: No such ID Found For Delete ")
	}
	return nil
}
