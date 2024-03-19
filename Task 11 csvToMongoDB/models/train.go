package models

type Train struct {
	Index     int    `bson:"index" json:"index"`
	TrainNo   int    `bson:"train_no" json:"train_no"`
	TrainName string `bson:"train_name" json:"train_name"`
	Starts    string `bson:"start" json:"starts"`
	Ends      string `bson:"ends" json:"ends"`
}
