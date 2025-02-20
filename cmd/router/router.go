package router

import (
	"country-rest-api/constants"
	handler "country-rest-api/handler"
	"log"
	"net/http"
	"os"
	"time"
)

// StartServer initializes the HTTP server, sets up the routes, and starts listening on the
// specified port.
func StartServer() {

	// Retrieves the port number from the environment variable "port"
	port := os.Getenv("port")
	if port == "" {
		// If not set, defaults to port 8080.
		log.Println("$port has not been set. Default: 8080")
		port = "8080"
	}

	// Creates a new HTTP request multiplexer (router) to handle incoming requests
	router := http.NewServeMux()

	// Registers various route handlers to the router
	router.HandleFunc(constants.DefaultPath, handler.EmptyHandler)
	router.HandleFunc(constants.InfoPath+constants.Iso2, handler.InfoHandler)
	router.HandleFunc(constants.PopulationPath+constants.Iso2, handler.PopulationHandler)
	router.HandleFunc(constants.StatusPath, handler.StatusHandler)

	// Logs when the server was started.
	constants.StartTime = time.Now()

	// Starts the HTTP server on the specified port and logs any fatal errors that occur.
	log.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
