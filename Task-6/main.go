package main

import (
	"crud_postgres/api"
	"crud_postgres/dbconn"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to Postgres DB and return db connection
	app, err := dbconn.NewPostgresDBConn()
	if err != nil {
		panic(err)
	}

	// Close DB connection when Program ends
	defer app.ClosePostgresConn()

	// Make new API struct and pass it to ServeMuxRouter function
	api := api.NewApi(app)

	serveMuxRouter(api)
}

func serveMuxRouter(myAPI *api.API) {

	// Get Mux New Router and Pass it to Initialize Router function
	muxRouter := mux.NewRouter()
	myAPI.InitializeRouters(muxRouter)

	// Starting server at 3452
	log.Println("server started at 3452")
	http.ListenAndServe(":3452", muxRouter)
}
