package grpcControllers

import (
	"Task-12/grpcService"
	"context"
	"database/sql"
	"log"
)

// Struct to Implement RouteStations interface in GRPC
type RouteStationServer struct {
	grpcService.UnimplementedRouteStationServiceServer
	DB *sql.DB
}

// Struct to Implement Route interface in GRPC
type RouteServer struct {
	grpcService.UnimplementedRouteServiceServer
	DB *sql.DB
}

// Get the Route Details by providing RouteId
func (r *RouteServer) GetRoute(ctx context.Context, in *grpcService.RouteRequest) (*grpcService.Route, error) {
	// Query for getting Info from RouteId
	res, err := r.DB.Query("SELECT * FROM transport.route where id = $1", in.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Store Data in struct
	route := grpcService.Route{}
	for res.Next() {
		res.Scan(&route.Id, &route.Name, &route.Status, &route.Source, &route.Destination)
	}

	return &route, nil
}

// Get all the Stations & Station Order come in Route given by Client
func (r *RouteServer) GetStationOrder(ctx context.Context, in *grpcService.RouteRequest) (*grpcService.RouteStationResponse, error) {
	// Query for getting Info from RouteStation table
	res, err := r.DB.Query("SELECT * FROM transport.routestations where route_id = $1", in.Id)
	if err != nil {
		return nil, err
	}

	// Store the Order details in struct and return to the client
	orders := []*grpcService.RouteStation{}
	for res.Next() {
		singleOrder := &grpcService.RouteStation{}
		res.Scan(&singleOrder.RouteId, &singleOrder.StationId, &singleOrder.StationOrder)
		orders = append(orders, singleOrder)
	}

	return &grpcService.RouteStationResponse{RouteStations: orders}, nil
}

// Provide the Station Details by StationId
func (r *RouteStationServer) GetAllStationOfRoute(ctx context.Context, in *grpcService.RouteRequest) (*grpcService.Station, error) {
	// Query to get the station by ID
	res, err := r.DB.Query("SELECT * FROM transport.station where id = $1", in.Id)
	if err != nil {
		return nil, err
	}

	// Store all Data in station struct and return to the client
	station := grpcService.Station{}
	for res.Next() {
		res.Scan(&station.Id, &station.Name, &station.Lat, &station.Long)
	}
	return &station, nil
}
