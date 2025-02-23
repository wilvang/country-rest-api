package service

import (
	"context"
	"country-rest-api/constants"
	"country-rest-api/util"
	"io"
	"log"
	"net/http"
	"time"
)

// IssuePostRequest sends a POST request to the specified URL with the given body.
// It uses the context from the incoming request and sets a timeout of 5 seconds.
// url: The URL to send the POST request to.
// body: The request body to be sent, encoded as JSON.
// r: The incoming HTTP request, used to derive the context.
// Returns the response body as a byte slice.
func IssuePostRequest(url string, body interface{}, r *http.Request) []byte {
	// Create a context with a timeout to ensure the request does not hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Encode the request body as JSON
	jsonBody := util.EncodeJSONBody(body)

	// Create a new HTTP request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, jsonBody)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
	}
	req.Header.Set("Content-Type", "application/json")

	return issueRequest(req)
}

// IssueGetRequest sends a GET request to the specified URL.
// It uses the context from the incoming request and sets a timeout of 5 seconds.
// url: The URL to send the GET request to.
// r: The incoming HTTP request, used to derive the context.
// Returns the response body as a byte slice.
func IssueGetRequest(url string, r *http.Request) []byte {
	// Create a context with a timeout to ensure the request does not hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Create a new HTTP request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
	}
	req.Header.Set("Content-Type", "application/json")

	return issueRequest(req)
}

// issueRequest sends the given HTTP request and returns the response body as a byte slice.
// req: The HTTP request to be sent.
// Returns the response body as a byte slice.
func issueRequest(req *http.Request) []byte {
	// Send the HTTP request and get the response
	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
	}
	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	// Read the response body
	body, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		log.Printf(constants.ErrorReadBody, err3)
	}

	return body
}
