package handler

import (
	"country-rest-api/constants"
	"country-rest-api/internal/service/info"
	"country-rest-api/internal/service/status"
	"country-rest-api/util"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// InfoHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handleInfoRequest(w, r)
		break
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handleInfoRequest processes GET requests to retrieve country information.
// It extracts the path parameter and query parameters, calls the service to get the country data,
// and returns the data in JSON format. If an error occurs, it returns an appropriate HTTP error response.
func handleInfoRequest(w http.ResponseWriter, r *http.Request) {

	// Extract path parameter
	param := strings.TrimPrefix(r.URL.Path, constants.InfoPath)
	if len(param) != 2 {
		http.Error(w, "Invalid path-parameter. Remember to use the iso2 for the desired country",
			http.StatusBadRequest)
		return
	}

	// Checks if the external APIs are running
	serverStatus := status.RequestStatusService(r)
	if serverStatus.CountriesNow != "200" || serverStatus.RestCountries != "200" {
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

func InfoPage(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("frontend/info.html")
	if err != nil {
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
