package info

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service/countriesNow"
	"country-rest-api/external/api/service/restCountries"
	"country-rest-api/models"
	"net/http"
	"strconv"
)

// RequestInfoService sends an HTTP GET request to the REST Countries API to retrieve information
// about a country specified by the 'param' parameter. It returns an Info struct with the decoded
// data or an error if the request or decoding fails.
func RequestInfoService(param string, limit string, r *http.Request) models.Info {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.InfoFilter

	// Send request to REST Countries API and retrieve country information
	countryResponse := restCountries.RequestInfo(url, r)
	// Send request to Countries Now API and retrieve city information
	cityResponse := countriesNow.RequestInfo(countryResponse.Name.Common, r)

	// Populate the Info struct with the retrieved data
	info := models.Info{
		Country:    countryResponse.Name.Common,
		Continents: countryResponse.Continents,
		Population: countryResponse.Population,
		Languages:  countryResponse.Languages,
		Borders:    countryResponse.Borders,
		Flag:       countryResponse.Flags.Png,
		Capitals:   countryResponse.Capital,
		Cities:     cityResponse[:10],
	}

	// Adjust the number of cities based on the 'limit' query parameter
	if lim, err := strconv.Atoi(limit); err == nil && lim >= 0 && lim <= len(cityResponse) {
		info.Cities = cityResponse[:lim]
	} else if limit != "" && lim < len(cityResponse) {
		info.Cities = cityResponse
	}

	return info
}
