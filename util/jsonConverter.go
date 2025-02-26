package util

import (
	"bytes"
	"country-rest-api/constants"
	"encoding/json"
	"log"
)

// DecodeJSONBody reads the response body and decodes it into the specified struct type.
func DecodeJSONBody(body []byte, model interface{}) {
	// Unmarshal the JSON data into the model
	err := json.Unmarshal(body, model)
	if err != nil {
		log.Printf(constants.ErrorDecodeJSON, err)
	}
}

// EncodeJSONBody reads the model and encodes it to JSON data.
func EncodeJSONBody(model interface{}) *bytes.Buffer {
	// Marshal the request data into JSON
	jsonData, err := json.Marshal(model)
	if err != nil {
		log.Printf(constants.ErrorEncodeJSON, err)
		return nil
	}
	return bytes.NewBuffer(jsonData)
}