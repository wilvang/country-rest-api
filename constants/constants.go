package constants

import "time"

// Path variables
const (
	// Iso2 is the placeholder used for inserting the two-letter country code in the URL.
	Iso2 = "{two_letter_country_code}"

	InfoFilter = "?fields=name,continents,population,languages,borders,flags,capital"
)

// StartTime  represents the time when the webserver started.
var StartTime time.Time

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
	// CountriesNowAPI endpoint for the Countries Now service.
	CountriesNowAPI = "http://129.241.150.113:3500/api/v0.1/"

	CountriesNowStatus = CountriesNowAPI + "countries/iso"

	// RESTCountriesAPI endpoint for the REST Countries service.
	RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/"

	RESTCountriesStatus = RESTCountriesAPI + "alpha/no?filter=capital"
)
