package constants

// Error messages
const (
	ErrorCloseBody      = "error closing response body: %v"
	ErrorReadBody       = "error reading response body: %v"
	ErrorCreateRequest  = "error in creating request: %v"
	ErrorResponse       = "error in response: %v"
	ErrorDecodeJSON     = "error decoding JSON: %v"
	ErrorEncodeJSON     = "error encoding JSON: %v"
	ErrorWritingJSON    = "error writing JSON file to response"
	ErrorWritingHTML    = "error writing HTML file to response"
	ErrorPrettyPrinting = "error during pretty printing"
	ErrorReadingHTML    = "error reading HTML file"
	ErrorPathParameter  = "error: invalid path-parameter"
	ErrorConnection     = "error: cannot connect to services"
	ErrorNotFound       = "error: request not found in the database, invalid iso2 code"
)
