package response

// PopulationResponse represents the response structure for population-related API requests.
type PopulationResponse struct {
	Error bool         `json:"error"`
	Msg   string       `json:"msg"`
	Data  DataResponse `json:"data"`
}

// DataResponse represents detailed population data for a specific country.
type DataResponse struct {
	Country          string            `json:"country"`
	Code             string            `json:"code"`
	ISO3             string            `json:"iso3"`
	PopulationCounts []PopulationCount `json:"populationCounts"`
}

// PopulationCount represents population data for a specific year.
type PopulationCount struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}
