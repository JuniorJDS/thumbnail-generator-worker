package api

import (
	"thumbnail-generator-worker/api/route"

	"github.com/gorilla/mux"
)

func HttpHandler() *mux.Router {
	router := mux.NewRouter()

	// define routes
	helloRoute := route.NewHelloWorld()
	router.HandleFunc("/item", helloRoute.GetHelloWorld).Methods("GET")

	return router
}
