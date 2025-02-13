package main

import (
	"country-rest-api/config"
	"country-rest-api/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	// Handle port assignment based on environment variable, or local override.
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handlers for all endpoints.
	http.HandleFunc(config.DefaultPath, handler.EmptyHandler)
	http.HandleFunc(config.InfoPath, handler.InfoHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
