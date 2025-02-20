package restCountries

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

	url = "https://countriesnow.space/api/v0.1/countries/iso"
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return fmt.Sprintf("Status Code: %d", http.StatusBadRequest)
	}

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return fmt.Sprintf("Status Code: %d", http.StatusInternalServerError)
	}
	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	return fmt.Sprintf("Status Code: %d", resp.StatusCode)
}
