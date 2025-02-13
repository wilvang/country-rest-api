package config

const (
	// DefaultPath for guiding the user to the correct path.
	DefaultPath = "/"

	// InfoPath for retrieving country information.
	InfoPath = "/countryinfo/v1/info/{:}"

	// PopulationPath for retrieving population records for a country.
	PopulationPath = "/countryinfo/v1/population/"

	// StatusPath for checking the uptime of the services.
	StatusPath = "/countryinfo/v1/status/"

	// CountriesNowAPI endpoint for the Countries Now service.
	CountriesNowAPI = "http://129.241.150.113:3500/api/v0.1/"

	// RESTCountriesAPI endpoint for the REST Countries service.
	RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/"
)
