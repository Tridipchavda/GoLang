package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to connect to MongoDB and return mongodb client instance
func ConnMongoDB() *mongo.Client {
	// set URI in options for localhost
	options := options.Client().ApplyURI("mongodb://localhost:27017")
	// Try to connect to MongoDB and handle error
	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		log.Println(err)
		return nil
	}
	// Check the connection with Ping and handle error
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Could not connect to MongoDB:", err)
		return nil
	}
	// Connect to Mongo and return client
	log.Println("Connected to MongoDB!")
	return client
}
