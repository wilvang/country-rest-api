package handler

import (
	"country-rest-api/internal/constants"
	"net/http"
	"strings"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
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
