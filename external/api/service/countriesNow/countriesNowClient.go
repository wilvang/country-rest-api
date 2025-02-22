package countriesNow

import (
	"context"
	"country-rest-api/constants"
	"country-rest-api/models/countriesNow/request"
	"country-rest-api/models/countriesNow/response"
	"country-rest-api/util"
	"io"
	"log"
	"net/http"
	"time"
)

// RequestInfo sends an HTTP POST request to the Countries Now API to retrieve city information
// for a specified country. It returns a slice of city names. If an error occurs during the request
// or decoding, it logs the error and returns an empty slice.
func RequestInfo(country string, r *http.Request) []string {

	// Create a context with a timeout to ensure the request does not hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	cityResponse := response.CityResponse{}

	// Encode the request body as JSON
	jsonBody := util.EncodeJSONBody(request.PostRequestBody{Name: country})

	// Create a new HTTP request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, constants.CountriesNowAPI+"countries/cities", jsonBody)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return make([]string, 0)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return make([]string, 0)
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
		return make([]string, 0)
	}

	// Decode the JSON response into the CityResponse struct
	util.DecodeJSONBody(body, &cityResponse)

	return cityResponse.Data
}
