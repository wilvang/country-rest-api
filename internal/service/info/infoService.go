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
// about a country specified by the 'param' parameter. It returns an Info struct with the decoded
// data or an error if the request or decoding fails.
func RequestInfoService(param string, limit string, r *http.Request) models.Info {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.InfoFilter

	countryResponse := restCountries.RequestInfo(url, r)
	cityResponse := countriesNow.RequestInfo(countryResponse.Name.Common, r)

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

	if lim, err := strconv.Atoi(limit); err == nil && lim >= 0 && lim <= len(cityResponse) {
		info.Cities = cityResponse[:lim]
	} else if lim < len(cityResponse) {
		info.Cities = cityResponse
	}

	return info
}
