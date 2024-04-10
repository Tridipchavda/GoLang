package main

import (
	dbconn "Task-12/dbConn"
	grpcHandle "Task-12/grpc-handle"
	httpRoutes "Task-12/http-handle"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Get the postgres sql instance
	postgresDB, err := dbconn.ConnectDB()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// Connect to GRPC server when HTTP server get started
	conn, err := grpc.Dial("localhost:7799", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Get Instance of App
	app := httpRoutes.NewApp(conn)
	// Closing Connection at end
	defer conn.Close()

	// Handle GRPC bindings and Serving
	grpcHandle.ConnectAndBind(postgresDB)
	// Handle HTTP Routes and Serving
	httpRoutes.InitHttpRouter(*app)
}
