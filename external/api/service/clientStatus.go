package service

import (
	"context"
	"country-rest-api/constants"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// RequestStatus sends an HTTP HEAD request to the specified URL to check the status of an external service.
// It returns the HTTP status code as a string. If an error occurs during the request, it logs the error
// and returns an appropriate HTTP status code.
func RequestStatus(url string, r *http.Request) string {

	// Create a context with a timeout to ensure the request does not hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Create a new HTTP request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return fmt.Sprint(http.StatusBadRequest)
	}

	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return fmt.Sprint(http.StatusInternalServerError)
	}
	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	// Return the HTTP status code as a string
	return fmt.Sprint(resp.StatusCode)
}
