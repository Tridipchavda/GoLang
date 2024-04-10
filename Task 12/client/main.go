package main

import (
	"Task-12/grpcService"
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:7799", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Create gRPC client Clients
	clientForRoutes := grpcService.NewRouteServiceClient(conn)
	clientForStations := grpcService.NewRouteStationServiceClient(conn)

	// Perform RPC calls
	routeRequest := &grpcService.RouteRequest{Id: 2}

	var route *grpcService.Route
	var stationOrder *grpcService.RouteStationResponse
	// Define WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrenctly calling Two grpc Services
	go func() {
		// Get route Info for given Id
		route, err = clientForRoutes.GetRoute(context.Background(), routeRequest)
		defer wg.Done()
		if err != nil {
			log.Fatalf("could not get station details: %v", err)
		}
	}()

	go func() {
		// Get Station Order Info for given Id
		stationOrder, err = clientForRoutes.GetStationOrder(context.Background(), routeRequest)
		defer wg.Done()
		if err != nil {
			log.Fatalf("could not get station details: %v", err)
		}
	}()
	// Waiting for all Data to Set
	wg.Wait()
	// Struct to store the all Station Response
	var stations []grpcService.Station

	log.Println(route)
	// Get all Station Info from the station_id get in station_order
	for _, v := range stationOrder.RouteStations {
		// Request for all Stations that come in the given Route
		routeRequest = &grpcService.RouteRequest{Id: v.StationId}
		station, _ := clientForStations.GetAllStationOfRoute(context.Background(), routeRequest)
		stations = append(stations, *station)
	}

	// Printing the Result
	for i := range stations {
		log.Printf("%v %v %v", stationOrder.RouteStations[i].StationOrder, route.Name, stations[i].Name)
	}

}
