package service

import (
	"bytes"
	"country-rest-api/internal/constants"
	"country-rest-api/internal/models"
	jsonutil "country-rest-api/internal/util/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// RequestInfoService sends an HTTP GET request to the REST Countries API to retrieve information
// about a country specified by the 'param' parameter. It returns a Country struct with the decoded
// data or an error if the request or decoding fails.
func RequestInfoService(param string, limit string) (models.Country, error) {
	filter := "?fields=name,continents,population,languages,borders,flags,capital"
	url := constants.RESTCountriesAPI + "alpha/" + param + filter

	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.Country{}, fmt.Errorf(constants.ErrorCreateRequest, err)
	}

	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		return models.Country{}, fmt.Errorf(constants.ErrorResponse, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(constants.ErrorCloseBody, err)
		}
	}(res.Body)

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		return models.Country{}, fmt.Errorf(constants.ErrorStatusCode, res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Country{}, fmt.Errorf("error reading response body: %v", err)
	}

	country := models.Country{}

	// Decode the JSON data into the Country struct
	err = jsonutil.DecodeJSONBody(io.NopCloser(bytes.NewReader(body)), &country)
	if err != nil {
		return country, err
	}

	// Updates the struct with the nested fields
	err = extractNestedFields(body, &country)
	if err != nil {
		return country, err
	}

	err = requestCities(&country, limit)
	if err != nil {
		return country, err
	}

	return country, nil
}

// extractNestedFields extracts the nested fields 'name' and 'flag' from the JSON response body.
func extractNestedFields(body []byte, country *models.Country) error {
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
func requestCities(country *models.Country, limit string) error {
	url := constants.CountriesNowAPI + "countries/cities"

	cityResponse := models.CityResponse{
		Error: true,
		Msg:   "",
		Data:  make([]string, 0),
	}

	// Create instance of content
	jsonBody, err := jsonutil.EncodeJSONBody(models.CityRequest{Name: country.Country})
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, url, jsonBody)
	if err != nil {
		return fmt.Errorf(constants.ErrorCreateRequest, err)
	}
	r.Header.Set("Content-Type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf(constants.ErrorResponse, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf(constants.ErrorCloseBody, err)
		}
	}(res.Body)

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		log.Printf("error: received status code %d", res.StatusCode)
		return fmt.Errorf(constants.ErrorStatusCode, res.StatusCode)
	}

	// Decode the JSON response body into the cityResponse model
	err = jsonutil.DecodeJSONBody(res.Body, &cityResponse)
	if err != nil {
		return err
	}

	// Slices the list of cities to the limit if provided
	if lim, err := strconv.Atoi(limit); err == nil && lim >= 0 && lim <= len(cityResponse.Data) {
		country.Cities = cityResponse.Data[:lim]
	} else {
		country.Cities = cityResponse.Data
	}

	return nil
}
