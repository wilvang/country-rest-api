package response

// CityResponse represents the response containing a list of cities for a specific country.
type CityResponse struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}
