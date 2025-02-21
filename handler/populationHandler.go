package handler

import (
	"country-rest-api/constants"
	"log"
	"net/http"
	"os"
	"strings"
)

// PopulationHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handlePopulationRequest(w, r)
		break
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handlePopulationRequest processes GET requests to retrieve the server status information.
// It calls the service to get the server status data and returns the data in JSON format.
// If an error occurs, it returns an appropriate HTTP error response.
func handlePopulationRequest(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter
	param := strings.TrimPrefix(r.URL.Path, constants.PopulationPath)

	// Extract query parameters
	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")

	// Use param and limit as needed
	response := "Received param: " + param + ", limit: " + limit
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func PopulationPage(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("frontend/population.html")
	if err != nil {
		log.Printf("Error writing HTML file to response: %v", err)
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML file content to the response
	_, err2 := w.Write(htmlFile)
	if err2 != nil {
		log.Printf("Error writing HTML file to response: %v", err2)
		http.Error(w, "Error writing HTML file to response", http.StatusInternalServerError)
	}
}
