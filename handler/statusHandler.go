package handler

import (
	"country-rest-api/internal/service/status"
	"encoding/json"
	"net/http"
)

// InfoHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handleGetRequest(w, r)
		break
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
	serverStatus := status.RequestStatusService(r)

	output, err := json.MarshalIndent(serverStatus, "", "  ")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	// Set the content type and serverStatus code before writing the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
