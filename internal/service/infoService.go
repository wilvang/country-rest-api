package service

import (
	"bytes"
	"country-rest-api/internal/constants"
	"country-rest-api/internal/models"
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
		return models.Country{}, fmt.Errorf("error in creating request: %v", err)
	}

	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		return models.Country{}, fmt.Errorf("error in response: %v", err)
	}
	defer res.Body.Close()

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		return models.Country{}, fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.Country{}, fmt.Errorf("error reading response body: %v", err)
	}

	// Initialize a Country struct to hold the decoded data
	country := models.Country{}

	// Decode the JSON data into the Country struct
	err = json.Unmarshal(body, &country)
	if err != nil {
		// If an error occurs during decoding, return a bad request status with the error message
		return models.Country{}, fmt.Errorf("error during decoding: %v", err)
	}

	// Updates the struct with the nested fields
	country.Country, country.Flag = extractNestedFields(body)

	cities, err := requestCities(&country, limit)
	country.Cities = cities

	return country, nil
}

// extractNestedFields extracts the nested fields 'name' and 'flag' from the JSON response body.
func extractNestedFields(body []byte) (string, string) {

	// Unmarshal JSON into a map
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Errorf("error during decoding: %v", err)
	}

	// Extract the nested values
	name := result["name"].(map[string]interface{})["common"].(string)
	flag := result["flags"].(map[string]interface{})["png"].(string)

	return name, flag
}

// requestCities sends an HTTP POST request to the CountriesNow API to retrieve a list of cities
// for the specified country. It returns a slice of city names or an error if the request or decoding fails.
func requestCities(country *models.Country, limit string) ([]string, error) {

	url := constants.CountriesNowAPI + "countries/cities"

	cityResponse := models.CityResponse{
		Error: true,
		Msg:   "",
		Data:  make([]string, 0),
	}

	// Create instance of content
	requestBody := models.CityRequest{Name: country.Country}

	// Marshal the request data into JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("error during decoding: %v", err)
		return cityResponse.Data, fmt.Errorf("error marshalling request data: %v", err)
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("error during decoding: %v", err)
		return cityResponse.Data, fmt.Errorf("error in creating request: %v", err)
	}
	r.Header.Set("Content-Type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	res, err := client.Do(r)
	if err != nil {
		log.Printf("error during decoding: %v", err)
		return cityResponse.Data, fmt.Errorf("error in response: %v", err)
	}
	defer res.Body.Close()

	// Check if the response status is not OK
	if res.StatusCode != http.StatusOK {
		log.Printf("error during decoding: %v", err)
		return cityResponse.Data, fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error during decoding: %v", err)
		return cityResponse.Data, fmt.Errorf("error reading response body: %v", err)
	}

	// Decodes the json response
	if err := json.Unmarshal(body, &cityResponse); err != nil {
		log.Printf("error during decoding: HERE %v", err)
		return cityResponse.Data, fmt.Errorf("error during decoding: %v", err)
	}

	// Slices the list of cities to the limit if provided
	if lim, err := strconv.Atoi(limit); err == nil && lim >= 0 {
		return cityResponse.Data[:lim], nil
	}

	// returns the complete list of cities
	return cityResponse.Data, nil
}
