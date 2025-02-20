package request

// PostRequestBody represents a request to retrieve info from CountriesNowAPI
type PostRequestBody struct {
	Name string `json:"country"`
}
