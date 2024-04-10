package httpRoutes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Struct to handle grpc Client Connection in http
type App struct {
	conn *grpc.ClientConn
}

// Function to Create and return instance of App
func NewApp(conn *grpc.ClientConn) *App {
	return &App{conn: conn}
}

// Function act as Gateway between HTTP client Request and GRPC Server
// Write Route Details
func (a *App) HandleGetRouteDetails(w http.ResponseWriter, r *http.Request) {
	// Call GRPC for route ask by user
	stringId, _ := mux.Vars(r)["i"]
	// Convert RouteId from string to int
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Function to call GRPC methods and get the data as "RouteDetails" Array
	details := callGRPCForData(uint32(id), a.conn)
	// Set Header as Json for JSON response
	w.Header().Set("Content-Type", "application/json")
	// Marshal the Data from struct and Handle Error for Same
	res, err := json.Marshal(details)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Provide JSON data for RouteDetails
	w.Write(res)
}

func InitHttpRouter(a App) {
	// Get Mux Router Instance
	m := mux.NewRouter()
	// Handle Route for getRouteDetails
	m.HandleFunc("/getRouteDetails/{i}", a.HandleGetRouteDetails)

	// Initializing server for HTTP
	log.Println("HTTP server listening on PORT 7788")
	if err := http.ListenAndServe(":7788", m); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
