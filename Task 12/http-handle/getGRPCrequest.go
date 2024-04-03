package httpRoutes

import (
	"Task-12/grpcService"
	"Task-12/models"
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
)

func callGRPCForData(id uint32, conn *grpc.ClientConn) []models.RouteDetails {
	// Set up a connection to the server.
	// Create gRPC client Clients
	clientForRoutes := grpcService.NewRouteServiceClient(conn)
	clientForStations := grpcService.NewRouteStationServiceClient(conn)

	// Perform RPC calls
	routeRequest := &grpcService.RouteRequest{Id: id}

	var err error
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

	// Printing the Result and store the result in routeDetails
	routeDetails := []models.RouteDetails{}
	for i := range stations {
		log.Printf("%v %v %v", stationOrder.RouteStations[i].StationOrder, route.Name, stations[i].Name)
		particularDetail := models.RouteDetails{
			Id:          stationOrder.RouteStations[i].StationOrder,
			RouteName:   route.Name,
			StationName: stations[i].Name,
			Lat:         stations[i].Lat,
			Long:        stations[i].Long,
		}
		routeDetails = append(routeDetails, particularDetail)
	}

	return routeDetails
}
