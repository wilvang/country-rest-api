package models

// CityRequest represents a request to retrieve cities for a specific country.
type CityRequest struct {
	// Name is the name of the country for which cities are being requested.
	Name string `json:"country"`
}

// CityResponse represents the response containing a list of cities for a specific country.
type CityResponse struct {
	// Error indicates whether there was an error processing the request.
	Error bool `json:"error"`

	// Msg contains a message related to the request or response.
	Msg string `json:"msg"`

	// Data is a list of city names in the specified country.
	Data []string `json:"data"`
}
