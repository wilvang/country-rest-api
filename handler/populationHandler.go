package handler

import (
	"country-rest-api/constants"
	"country-rest-api/internal/service/population"
	"country-rest-api/internal/service/status"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// PopulationHandler handles HTTP requests and routes them to the appropriate handler function
// based on the HTTP method. If the method is not supported, it returns a
// "Not Implemented" status.
// w: The HTTP response writer.
// r: The HTTP request.
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// Selects the appropriate function for the HTTP request.
	switch r.Method {
	case http.MethodGet:
		handlePopulationRequest(w, r)
	default:
		http.Error(w, "REST Method '"+r.Method+"' not supported. Currently only '"+
			http.MethodGet+"' is supported.", http.StatusNotImplemented)
		return
	}
}

// handlePopulationRequest processes GET requests to retrieve population information.
// It calls the service to get the population data and returns the data in JSON format.
// If an error occurs, it returns an appropriate HTTP error response.
// w: The HTTP response writer.
// r: The HTTP request.
func handlePopulationRequest(w http.ResponseWriter, r *http.Request) {
	// Extract path parameter
	param := strings.TrimPrefix(r.URL.Path, constants.PopulationPath)
	if len(param) != 2 {
		http.Error(w, constants.ErrorPathParameter, http.StatusBadRequest)
		return
	}

	// Checks if the external APIs are running
	serverStatus := status.RequestStatusService(r)
	if serverStatus.CountriesNow != "200" || serverStatus.RestCountries != "200" {
		http.Error(w, constants.ErrorConnection, http.StatusInternalServerError)
		return
	}

	// Extract query parameters
	queryParams := r.URL.Query()
	limit := queryParams.Get("limit")

	// Call the service to get the country information
	populationInfo := population.RequestPopulationService(param, limit, r)

	// Pretty-print the JSON response
	output, err := json.MarshalIndent(populationInfo, "", "  ")
	if err != nil {
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

// PopulationPage serves the population HTML page.
// w: The HTTP response writer.
// r: The HTTP request.
func PopulationPage(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlFile, err := os.ReadFile("frontend/population.html")
	if err != nil {
		log.Printf(constants.ErrorReadingHTML+" %v", err)
		http.Error(w, constants.ErrorReadingHTML, http.StatusInternalServerError)
		return
	}

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML file content to the response
	_, err2 := w.Write(htmlFile)
	if err2 != nil {
		log.Printf(constants.ErrorWritingHTML+" %v", err2)
		http.Error(w, constants.ErrorWritingHTML, http.StatusInternalServerError)
	}
}
