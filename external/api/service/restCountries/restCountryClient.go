package restCountries

import (
	"country-rest-api/external/api/service"
	"country-rest-api/models/restCounties/response"
	"country-rest-api/util"
	"net/http"
)

// RequestInfo sends an HTTP GET request to the specified URL to retrieve country information.
// It returns a CountryResponse struct with the decoded data. If an error occurs during the request
// or decoding, it logs the error and returns an empty CountryResponse struct.
// url: The URL to send the GET request to.
// r: The incoming HTTP request, used to derive the context.
// Returns a CountryResponse struct containing the country information.
func RequestInfo(url string, r *http.Request) response.CountryResponse {
	resp := response.CountryResponse{}

	// Send the GET request and get the response body
	body := service.IssueGetRequest(url, r)

	// Decode the JSON response into the CountryResponse struct
	util.DecodeJSONBody(body, &resp)

	return resp
}

// RequestName sends an HTTP GET request to the specified URL to retrieve country name information.
// It returns a NameResponse struct with the decoded data. If an error occurs during the request
// or decoding, it logs the error and returns an empty NameResponse struct.
// url: The URL to send the GET request to.
// r: The incoming HTTP request, used to derive the context.
// Returns a NameResponse struct containing the country name information.
func RequestName(url string, r *http.Request) response.NameResponse {
	resp := response.NameResponse{}

	// Send the GET request and get the response body
	body := service.IssueGetRequest(url, r)

	// Decode the JSON response into the NameResponse struct
	util.DecodeJSONBody(body, &resp)

	return resp
}
