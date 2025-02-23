package response

import "country-rest-api/models"

// PopulationResponse represents the response structure for population-related API requests.
type PopulationResponse struct {
	Error bool         `json:"error"`
	Msg   string       `json:"msg"`
	Data  DataResponse `json:"data"`
}

// DataResponse represents detailed population data for a specific country.
type DataResponse struct {
	Country          string                  `json:"country"`
	Code             string                  `json:"code"`
	ISO3             string                  `json:"iso3"`
	PopulationCounts []models.PopulationData `json:"populationCounts"`
}
