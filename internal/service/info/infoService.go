package info

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service/countriesNow"
	"country-rest-api/external/api/service/restCountries"
	"country-rest-api/internal/models"
	"net/http"
	"strconv"
)

// RequestInfoService sends an HTTP GET request to the REST Countries API to retrieve information
// about a country specified by the 'param' parameter. It returns a Info struct with the decoded
// data or an error if the request or decoding fails.
func RequestInfoService(param string, limit string, r *http.Request) models.Info {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.InfoFilter

	countryResponse := restCountries.RequestInfo(url, r)

	info := models.Info{
		Country:    countryResponse.Name.Common,
		Continents: countryResponse.Continents,
		Population: countryResponse.Population,
		Languages:  countryResponse.Languages,
		Borders:    countryResponse.Borders,
		Flag:       countryResponse.Flags.Png,
		Capitals:   countryResponse.Capital,
		Cities:     nil,
	}

	cityResponse := countriesNow.RequestInfo(info.Country, r)

	if lim, err5 := strconv.Atoi(limit); err5 == nil && lim >= 0 && lim <= len(cityResponse) {
		info.Cities = cityResponse[:lim]
	} else {
		info.Cities = cityResponse
	}

	return info
}
