package main

import (
	"context"
	"log"

	"github.com/Tridipchavda/FiberWithMongoDB/apis"
	"github.com/Tridipchavda/FiberWithMongoDB/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to MongoDB
	mongoDB, err := configs.NewMongoDB()
	log.Println(mongoDB)

	// Handlinh the Error
	if err != nil {
		log.Println("Database Connection Error")
		return
	}
	// Create New Fiber App
	app := fiber.New()

	// Create new App passing MongoDB to it and Initalize Server with Fiber
	routes := apis.NewApp(mongoDB)
	routes.InitializeRoutes(app)

	// Disconnect from MongoDB
	if err := mongoDB.Disconnect(context.Background()); err != nil {
		panic(err)
	}

}
