package jsonutil

import (
	"bytes"
	"country-rest-api/constants"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

// DecodeJSONBody reads the response body and decodes it into the specified struct type.
func DecodeJSONBody(body io.ReadCloser, model interface{}) error {
	// Read the response body
	data, err := io.ReadAll(body)
	if err != nil {
		log.Printf(constants.ErrorReadBody, err)
		return fmt.Errorf(constants.ErrorReadBody, err)
	}

	// Unmarshal the JSON data into the model
	err = json.Unmarshal(data, model)
	if err != nil {
		log.Printf(constants.ErrorDecodeJSON, err)
		return fmt.Errorf(constants.ErrorDecodeJSON, err)
	}

	return nil
}

// EncodeJSONBody reads the model and encodes it to JSON data.
func EncodeJSONBody(model interface{}) (*bytes.Buffer, error) {
	// Marshal the request data into JSON
	jsonData, err := json.Marshal(model)
	if err != nil {
		log.Printf(constants.ErrorEncodeJSON, err)
		return nil, fmt.Errorf(constants.ErrorEncodeJSON, err)
	}
	return bytes.NewBuffer(jsonData), nil
}
