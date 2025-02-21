package handler

import (
	"country-rest-api/constants"
	"country-rest-api/internal/service/info"
	"country-rest-api/internal/service/status"
	"country-rest-api/util"
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
		handleStatusRequest(w, r)
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
func handleStatusRequest(w http.ResponseWriter, r *http.Request) {

	// Extract path parameter
	param := strings.TrimPrefix(r.URL.Path, constants.InfoPath)
	if len(param) != 2 {
		http.Error(w, "Invalid path-parameter. Remember to use the iso2 for the desired country",
			http.StatusBadRequest)
		return
	}

	// Checks if the external APIs are running
	serviceStatus := status.RequestStatusService(r)
	if serviceStatus.CountriesNow != "200" || serviceStatus.RestCountries != "200" {
		http.Error(w, "Cannot connect to the services", http.StatusInternalServerError)
		return
	}

	// Extract query parameters
	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")

	// Call the service to get the country information
	country := info.RequestInfoService(param, limit, r)
	if util.IsEmpty(country) {
		http.Error(w, "Error during fetching of data", http.StatusInternalServerError)
		return
	}

	// Pretty-print the JSON response
	output, err2 := json.MarshalIndent(country, "", "  ")
	if err2 != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}

	// Set the content type and status code before writing the response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
