package main

import (
	"fmt"
	"log"
	"net/http"

	"around/backend"
	"around/handler"
)

func main() {
	
	fmt.Println("started-service")
	
	// Initialize the backends
	backend.InitElasticsearchBackend()
	
	// Initialize GCS Backend
	backend.InitGCSBackend()
	
	
	// Start the HTTP server with the initialized router
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
