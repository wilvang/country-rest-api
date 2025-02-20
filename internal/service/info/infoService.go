package info

import (
	"bytes"
	"country-rest-api/constants"
	"country-rest-api/external/models/countriesNow/request"
	"country-rest-api/external/models/countriesNow/response"
	external "country-rest-api/internal/models"
	"country-rest-api/util/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// RequestInfoService sends an HTTP GET request to the REST Countries API to retrieve information
// about a country specified by the 'param' parameter. It returns an Info struct with the decoded
// data or an error if the request or decoding fails.
func RequestInfoService(param string, limit string) (external.Info, error) {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.InfoFilter

	r, err0 := http.NewRequest(http.MethodGet, url, nil)
	if err0 != nil {
		return external.Info{}, fmt.Errorf(constants.ErrorCreateRequest, err0)
	}

	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err1 := client.Do(r)
	if err1 != nil {
		return external.Info{}, fmt.Errorf(constants.ErrorResponse, err1)
	}

	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			log.Printf(constants.ErrorCloseBody, err2)
		}
	}(res.Body)

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		return external.Info{}, fmt.Errorf(constants.ErrorStatusCode, res.StatusCode)
	}

	// Read the response body
	body, err3 := io.ReadAll(res.Body)
	if err3 != nil {
		return external.Info{}, fmt.Errorf("error reading response body: %v", err3)
	}

	country := external.Info{}

	// Decode the JSON data into the Info struct
	err4 := jsonutil.DecodeJSONBody(io.NopCloser(bytes.NewReader(body)), &country)
	if err4 != nil {
		return country, err4
	}

	// Updates the struct with the nested fields
	err5 := extractNestedFields(body, &country)
	if err5 != nil {
		return country, err5
	}

	err6 := requestCities(&country, limit)
	if err6 != nil {
		return country, err6
	}

	return country, nil
}

// extractNestedFields extracts the nested fields 'name' and 'flag' from the JSON response body.
func extractNestedFields(body []byte, country *external.Info) error {
	// Unmarshal JSON into a map
	var result map[string]interface{}
	err := json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Extract the nested values
	country.Country = result["name"].(map[string]interface{})["common"].(string)
	country.Flag = result["flags"].(map[string]interface{})["png"].(string)

	return nil
}

// requestCities sends an HTTP POST request to the CountriesNow API to retrieve a list of cities
// for the specified country. It returns a slice of city names or an error if the request or decoding fails.
func requestCities(country *external.Info, limit string) error {
	url := constants.CountriesNowAPI + "countries/cities"

	cityResponse := response.CityResponse{
		Error: true,
		Msg:   "",
		Data:  make([]string, 0),
	}

	// Create instance of content
	jsonBody, err0 := jsonutil.EncodeJSONBody(request.PostRequestBody{Name: country.Country})
	if err0 != nil {
		return err0
	}

	r, err1 := http.NewRequest(http.MethodPost, url, jsonBody)
	if err1 != nil {
		return fmt.Errorf(constants.ErrorCreateRequest, err1)
	}
	r.Header.Set("Content-Type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err2 := client.Do(r)
	if err2 != nil {
		return fmt.Errorf(constants.ErrorResponse, err2)
	}

	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(res.Body)

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		log.Printf("error: received status code %d", res.StatusCode)
		return fmt.Errorf(constants.ErrorStatusCode, res.StatusCode)
	}

	// Decode the JSON response body into the cityResponse model
	err4 := jsonutil.DecodeJSONBody(res.Body, &cityResponse)
	if err4 != nil {
		return err4
	}

	// Slices the list of cities to the limit if provided
	if lim, err5 := strconv.Atoi(limit); err5 == nil && lim >= 0 && lim <= len(cityResponse.Data) {
		country.Cities = cityResponse.Data[:lim]
	} else {
		country.Cities = cityResponse.Data
	}

	return nil
}
