package restCountries

import (
	"context"
	"country-rest-api/constants"
	"country-rest-api/external/models/restCounties/response"
	"country-rest-api/util"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestInfo(url string, r *http.Request) response.CountryResponse {
	countryResponse := response.CountryResponse{}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return countryResponse
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return countryResponse
	}

	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	body, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		log.Printf(constants.ErrorReadBody, err3)
		return countryResponse
	}

	util.DecodeJSONBody(body, &countryResponse)

	return countryResponse
}
