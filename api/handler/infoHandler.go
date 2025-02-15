package handler

import (
	"country-rest-api/internal/constants"
	"country-rest-api/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

// InfoHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handleGetRequest processes GET requests to retrieve country information.
// It extracts the path parameter and query parameters, calls the service to get the country data,
// and returns the data in JSON format. If an error occurs, it returns an appropriate HTTP error response.
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter
	param := strings.TrimPrefix(r.URL.Path, constants.InfoPath)

	// Extract query parameters
	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")

	// Call the service to get the country information
	country, err := service.RequestInfoService(param, limit)
	if err != nil {
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Pretty-print the JSON response
	output, err := json.MarshalIndent(country, "", "  ")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	// Set the content type and status code before writing the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
