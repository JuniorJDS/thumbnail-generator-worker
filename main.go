package main

import (
	"log"
	"thumbnail-generator-worker/api"
	"thumbnail-generator-worker/api/app"
)

func main() {
	handler := api.HttpHandler()

	err := app.Start("5000", handler)
	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}
