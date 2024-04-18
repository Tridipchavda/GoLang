package main

import (
	"Task-14/router"
	"net/http"
)

func main() {
	// Create Mux Router and Bind it to Server listening at port 4321
	muxRouter := router.InitalizeRouter()
	http.ListenAndServe(":4321", muxRouter)
}
