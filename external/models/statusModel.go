package models

// Status represents the status information of the API services.
type Status struct {
	CountriesNow  string `json:"countriesnowapi"`
	RestCountries string `json:"restcountriesapi"`
	Version       string `json:"version"`
	Uptime        int    `json:"uptime"`
}
