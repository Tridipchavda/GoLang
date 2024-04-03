# GRPC Project

In this project 3 services are build on gRPC using by one HTTP GET request.

## Project Structure

- <b>/protobuf</b> Folder will be having all the proto files and their golang code will be stored inside <b>/grpcService</b> folder.

## Prerequisites
Before running the API, make sure you have the following installed:

- GoLang
- Postgres
- Mux Router
- protoc

## API EndPoints

### 1. Get All Station In Particular Routes 

#### REST API :  /getRouteDetails/ ( Note :- All three gRPC services are used by this one HTTP request only )

| Method    | Route| Description       |
| ----------| -----| ----------------- |
| GET       | /getRouteDetails    | Get All Station In Particular Routes by combining three datasets from different gRPC service    |

#### GRPC Services

```
service RouteService {
    rpc GetRoute (RouteRequest) returns (Route) {}
    rpc GetStationOrder (RouteRequest) returns (RouteStationResponse) {}
}

service RouteStationService {
    rpc GetAllStationOfRoute (RouteRequest) returns (Station) {}
}
```

## Running The Project

1. Clone the repository.
2. If Changing the protobuf files , run given Command in terminal 
```
make proto
```
OR
```
protoc --go_out=. --go-grpc_out=. ./protobuf/*.proto && cd ./grpcService && go get
```
3. Change the postgresurl from ./dbConn/initDB.go for Function ConnectDB()
4. go run main.go (Initalize GRPC server with HTTP/2 at 7799 as well as REST API SERVER AT PORT 7788 
5. In Client Folder , run go run main.go for Testing the apis with gRPC to gRPC calls
