package main

import (
	"country-rest-api/cmd/router"
)

// main is the entry point of the application. It starts the HTTP server by calling the StartServer
// function from the router package.
func main() {
	router.StartServer()
}
