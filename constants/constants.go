package constants

const (
	// PathValue represents the key for a two-letter country code in a URL path.
	PathValue = "two_letter_country_code"

	// PathParam represents the query parameter key for limiting the number of results.
	PathParam = "limit"

	// CountryCodePlaceholder is the placeholder used for inserting the two-letter country code in the URL.
	CountryCodePlaceholder = "{two_letter_country_code}"

	// LimitByCity represents a query parameter to limit the results by city.
	LimitByCity = "?limit={limit}"

	// LimitByYear represents a query parameter to limit the results by a range of years (startYear-endYear).
	LimitByYear = "{?limit={:startYear-endYear}}"
)

const (
	// DefaultPath for guiding the user to the correct path.
	DefaultPath = "/"

	// InfoPath for retrieving country information.
	InfoPath = "/countryinfo/v1/info/"

	// PopulationPath for retrieving population records for a country.
	PopulationPath = "/countryinfo/v1/population/"

	// StatusPath for checking the uptime of the services.
	StatusPath = "/countryinfo/v1/status/"
)

const (
	// CountriesNowAPI endpoint for the Countries Now service.
	CountriesNowAPI = "http://129.241.150.113:3500/api/v0.1/"

	// RESTCountriesAPI endpoint for the REST Countries service.
	RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/"
)

const (
	ErrorStatusCode    = "error: received status code %d"
	ErrorCloseBody     = "error closing response body: %v"
	ErrorReadBody      = "error reading response body: %v"
	ErrorDecodeJSON    = "error decoding JSON: %v"
	ErrorEncodeJSON    = "error encoding JSON: %v"
	ErrorCreateRequest = "error in creating request: %v"
	ErrorResponse      = "error in response: %v"
)
