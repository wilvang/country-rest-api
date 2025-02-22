package restCountries

import (
	"context"
	"country-rest-api/constants"
	"country-rest-api/models/restCounties/response"
	"country-rest-api/util"
	"io"
	"log"
	"net/http"
	"time"
)

// RequestInfo sends an HTTP GET request to the specified URL to retrieve country information.
// It returns a CountryResponse struct with the decoded data. If an error occurs during the request
// or decoding, it logs the error and returns an empty CountryResponse struct.
func RequestInfo(url string, r *http.Request) response.CountryResponse {
	countryResponse := response.CountryResponse{}

	// Create a context with a timeout to ensure the request does not hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Create a new HTTP request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return countryResponse
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return countryResponse
	}
	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	// Read the response body
	body, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		log.Printf(constants.ErrorReadBody, err3)
		return countryResponse
	}

	// Decode the JSON response into the CountryResponse struct
	util.DecodeJSONBody(body, &countryResponse)

	return countryResponse
}
