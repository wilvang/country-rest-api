package constants

// Error messages
const (
	ErrorCloseBody      = "error closing response body: %v"
	ErrorReadBody       = "error reading response body: %v"
	ErrorCreateRequest  = "error in creating request: %v"
	ErrorResponse       = "error in response: %v"
	ErrorDecodeJSON     = "error decoding JSON: %v"
	ErrorEncodeJSON     = "error encoding JSON: %v"
	ErrorWritingJSON    = "Error writing JSON file to response"
	ErrorWritingHTML    = "Error writing HTML file to response"
	ErrorPrettyPrinting = "Error during pretty printing"
	ErrorReadingHTML    = "Error reading HTML file"
	ErrorPathParameter  = "Invalid path-parameter! Remember to use the iso2 for the desired country"
	ErrorConnection     = "Cannot connect to services"
)
