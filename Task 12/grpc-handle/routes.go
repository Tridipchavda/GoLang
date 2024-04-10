package grpcHandle

import (
	"Task-12/grpcControllers"
	"Task-12/grpcService"
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Function to initialize the GRPC Server , add Services and Bind to TCP connection with HTTP/2
func ConnectAndBind(postgresDB *sql.DB) {
	// Get New GRPC Server
	RouteGRPCServer := grpc.NewServer()

	// Add / Register Service for RouteStation and pass DB to it
	routeStationServer := &grpcControllers.RouteStationServer{DB: postgresDB}
	grpcService.RegisterRouteStationServiceServer(RouteGRPCServer, routeStationServer)

	// Add / Register Service for Route and pass DB to it
	routeServer := &grpcControllers.RouteServer{DB: postgresDB}
	grpcService.RegisterRouteServiceServer(RouteGRPCServer, routeServer)

	// Start TCP with HTTP/2 Server at PORT 7799
	lis, err := net.Listen("tcp", ":7799")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Bind the GRPC server request to HTTP port
	go func() {
		log.Println("RouteStation gRPC service bind on PORT 7799")
		if err := RouteGRPCServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
}
