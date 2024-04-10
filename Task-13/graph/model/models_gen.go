package model

type Mutation struct {
}

type NewTrain struct {
	Index     int    `json:"index" bson:"index"`
	TrainNo   int    `json:"train_no" bson:"train_no"`
	TrainName string `json:"train_name" bson:"train_name"`
	Start     string `json:"start" bson:"start"`
	Ends      string `json:"ends" bson:"ends"`
}

type Query struct {
}

type Train struct {
	ID                  string    `json:"_id" bson:"_id"`
	Index               int       `json:"index" bson:"index"`
	TrainNo             int       `json:"train_no" bson:"train_no"`
	TrainName           string    `json:"train_name" bson:"train_name"`
	Start               string    `json:"start" bson:"start"`
	Ends                string    `json:"ends" bson:"ends"`
	TestTrainContainers []*string `json:"test_train_containers" bson:"test_train_containers"`
	Gettrains           []*Train  `json:"gettrains" bson:"gettrains"`
}

type CreateTrain struct {
	Index     int    `json:"index" bson:"index"`
	TrainNo   int    `json:"train_no" bson:"train_no"`
	TrainName string `json:"train_name" bson:"train_name"`
	Start     string `json:"start" bson:"start"`
	Ends      string `json:"ends" bson:"ends"`
}

type DeleteTrainInfo struct {
	DeleteTrainCount int `json:"deleteTrainCount"`
}

type UpdateTrain struct {
	Index     int    `json:"index" bson:"index"`
	TrainNo   int    `json:"train_no" bson:"train_no"`
	TrainName string `json:"train_name" bson:"train_name"`
	Start     string `json:"start" bson:"start"`
	Ends      string `json:"ends" bson:"ends"`
}
