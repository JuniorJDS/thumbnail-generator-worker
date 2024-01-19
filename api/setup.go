package api

import (
	"thumbnail-generator-worker/api/route"
	"thumbnail-generator-worker/api/route/endpoints"

	"github.com/gorilla/mux"
)

func HttpHandler() *mux.Router {
	router := mux.NewRouter()

	// define routes
	helloRoute := route.NewHelloWorld()
	router.HandleFunc("/", helloRoute.GetHelloWorld).Methods("GET")

	generateRoute := endpoints.NewThumbnailGeneratorRoute()
	router.HandleFunc("/", generateRoute.Generate).Methods("POST")

	return router
}
