package constants

import "time"

// StartTime  represents the time when the webserver started.
var StartTime time.Time

// Path variables
const (
	// Iso2 is the placeholder used for inserting the two-letter country code in the URL.
	Iso2 = "{two_letter_country_code}"

	// InfoFilter filters to only receive the necessary fields from the RESTCountriesAPI
	InfoFilter = "?fields=name,continents,population,languages,borders,flags,capital"

	CountryFilter = "?fields=name"
)

// Internal endpoints
const (
	// RootPath is the root endpoint of the API.
	RootPath = "/"

	// Version is the current version of the API.
	Version = "v1"

	// DefaultPath for guiding the user to the correct path.
	DefaultPath = RootPath + "countryinfo/" + Version

	// InfoPath for retrieving country information.
	InfoPath = DefaultPath + "/info/"

	// PopulationPath for retrieving population records for a country.
	PopulationPath = DefaultPath + "/population/"

	// StatusPath for checking the uptime of the services.
	StatusPath = DefaultPath + "/status/"
)

// External endpoints
const (
	// CountriesNowAPI is the endpoint for the Countries Now service.
	CountriesNowAPI = "http://129.241.150.113:3500/api/v0.1/"

	// CountriesNowStatus is the endpoint for retrieving country ISO codes from the Countries Now service.
	CountriesNowStatus = CountriesNowAPI + "countries/iso"

	CountriesNowInfo = CountriesNowAPI + "countries/cities"

	CountriesNowPopulation = CountriesNowAPI + "countries/population"

	// RESTCountriesAPI is the endpoint for the REST Countries service.
	RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/"

	// RESTCountriesStatus is the endpoint for retrieving the capital of Norway from the REST Countries service.
	RESTCountriesStatus = RESTCountriesAPI + "alpha/no?filter=capital"
)
