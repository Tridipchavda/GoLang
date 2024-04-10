package connMongoDB

import (
	"context"
	"graphQL/graph/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client          *mongo.Client
	trainCollection *mongo.Collection
}

func ConnectToMongoDB() *DB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	log.Println("Connected to MongoDB ")
	return &DB{client: client, trainCollection: client.Database("Train").Collection("Train")}
}


func (db *DB) Train(trainNo int) *model.Train {
	var train model.Train
	err := db.trainCollection.FindOne(context.TODO(), bson.M{"train_no": trainNo}).Decode(&train)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println(train)
	return &train
}

func (db *DB) Trains() []*model.Train {
	var trains []*model.Train
	cur, err := db.trainCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	for cur.Next(context.Background()) {
		var train model.Train
		cur.Decode(&train)
		trains = append(trains, &train)
	}
	return trains
}

func (db *DB) CreateTrain(input *model.NewTrain) *model.Train {

	res, err := db.trainCollection.InsertOne(context.TODO(), input)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println(res)
	return &model.Train{TrainName: "Got It"}
}

func (db *DB) DeleteTrain(trainNo int) *model.DeleteTrainInfo {

	res, err := db.trainCollection.DeleteMany(context.TODO(), bson.M{"train_no": trainNo})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	log.Println("Delete Count :", res)
	return &model.DeleteTrainInfo{DeleteTrainCount: int(res.DeletedCount)}
}

func (db *DB) UpdateTrain(trainNo int, input *model.NewTrain) *model.Train {

	res, err := db.trainCollection.UpdateOne(context.TODO(), bson.M{"train_no": trainNo},bson.M{"$set":input})
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println("Update Count :", res)
	return &model.Train{TrainName: input.TrainName,TrainNo: input.TrainNo,
		Start: input.Start,Ends: input.Ends,Index: input.Index}
}
