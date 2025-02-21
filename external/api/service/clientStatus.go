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

func RequestStatus(url string, r *http.Request) string {

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return fmt.Sprint(http.StatusBadRequest)
	}

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

	return fmt.Sprint(resp.StatusCode)
}
