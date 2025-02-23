package handler

import (
	"country-rest-api/constants"
	"country-rest-api/internal/service/status"
	"encoding/json"
	"log"
	"net/http"
)

// StatusHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handleStatusRequest(w, r)
		break
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handleStatusRequest processes GET requests to retrieve the server status information.
// It calls the service to get the server status data and returns the data in JSON format.
// If an error occurs, it returns an appropriate HTTP error response.
func handleStatusRequest(w http.ResponseWriter, r *http.Request) {
	// Retrieve the server status information
	serverStatus := status.RequestStatusService(r)

	// Pretty-print the JSON response
	output, err := json.MarshalIndent(serverStatus, "", "  ")
	if err != nil {
		log.Printf(constants.ErrorPrettyPrinting)
		http.Error(w, constants.ErrorPrettyPrinting, http.StatusInternalServerError)
		return
	}

	// Set the content type and status code before writing the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write(output)
	if err2 != nil {
		log.Printf(constants.ErrorWritingJSON+" %v", err2)
		http.Error(w, constants.ErrorWritingJSON, http.StatusInternalServerError)
	}
}
