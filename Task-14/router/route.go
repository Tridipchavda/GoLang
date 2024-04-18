package router

import (
	"github.com/gorilla/mux"
)

type Search []interface{}

func InitalizeRouter() *mux.Router {
	// Get new Router from mux
	muxRouter := mux.NewRouter()
	// Handling Routes for Query and Content of wikipedia
	muxRouter.HandleFunc("/search/{query}", HandleSearch)
	muxRouter.HandleFunc("/getContent/{title}", HandleWikiPediaContent)
	return muxRouter
}
