package handler

import (
	"country-rest-api/models"
	"encoding/json"
	"net/http"
)

// InfoHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func InfoHandler(w http.ResponseWriter, r *http.Request) {

	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodPost:
		handlePostRequest(w, r)
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' and '"+http.MethodPost+"' are supported.", http.StatusNotImplemented)
		return
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {

	// Create a new JSON decoder for the request body
	decoder := json.NewDecoder(r.Body)

	// Disallow unknown fields to ensure the JSON strictly matches the Country struct
	decoder.DisallowUnknownFields()

	// Initialize a Country struct to hold the decoded data
	country := models.Country{}

	// Decode the JSON data into the Country struct
	err := decoder.Decode(&country)
	if err != nil {
		// If an error occurs during decoding, return a bad request status with the error message
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}
}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	//TODO: Create implementation for GET requests.
}
