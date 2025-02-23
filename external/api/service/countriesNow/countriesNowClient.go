package countriesNow

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service"
	"country-rest-api/models/countriesNow/request"
	"country-rest-api/models/countriesNow/response"
	"country-rest-api/util"
	"net/http"
)

// RequestInfo sends an HTTP POST request to the Countries Now API to retrieve city information
// for a specified country. It returns a slice of city names. If an error occurs during the request
// or decoding, it logs the error and returns an empty slice.
// country: The name of the country for which to retrieve city information.
// r: The incoming HTTP request, used to derive the context.
// Returns a slice of city names.
func RequestInfo(country string, r *http.Request) []string {
	resp := response.CityResponse{}

	// Send the POST request to the Countries Now API
	body := service.IssuePostRequest(constants.CountriesNowInfo, request.PostRequestBody{Name: country}, r)

	// Decode the JSON response into the CityResponse struct
	util.DecodeJSONBody(body, &resp)

	return resp.Data
}

// RequestPopulation sends an HTTP POST request to the Countries Now API to retrieve population information
// for a specified country. It returns a PopulationResponse struct with the decoded data. If an error occurs
// during the request or decoding, it logs the error and returns an empty PopulationResponse struct.
// country: The name of the country for which to retrieve population information.
// r: The incoming HTTP request, used to derive the context.
// Returns a PopulationResponse struct containing the population information.
func RequestPopulation(country string, r *http.Request) response.PopulationResponse {
	resp := response.PopulationResponse{}

	// Send the POST request to the Countries Now API
	body := service.IssuePostRequest(constants.CountriesNowPopulation, request.PostRequestBody{Name: country}, r)

	// Decode the JSON response into the PopulationResponse struct
	util.DecodeJSONBody(body, &resp)

	return resp
}
