package utils

import (
	"context"
	"csvToMongo/models"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

// Function to insert the Data in MongoDB from csv file
func InsertDataFromCSV(fileName string, mongo *mongo.Client) error {
	// Open file given in the flag
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	_, _ = reader.Read()
	// Insert Data line by line accroding to CSV till end
	for {
		record, err := reader.Read()
		// If EOF occur return the function
		if err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		// convert index to integer
		index, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		// convert train_no to integer
		train_no, err := strconv.Atoi(record[1])
		if err != nil {
			return err
		}
		// Add data to mongoDB with collection Train in Database Train
		_, err = mongo.Database("Train").Collection("Train").InsertOne(
			context.Background(),
			models.Train{
				Index:     index,
				TrainNo:   train_no,
				TrainName: record[2],
				Starts:    record[3],
				Ends:      record[4],
			},
		)
		// Handle if any error occur between reading line by line
		if err != nil {
			return err
		}
	}
}
