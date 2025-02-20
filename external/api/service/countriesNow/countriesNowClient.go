package countriesNow

import (
	"context"
	"country-rest-api/constants"
	"country-rest-api/external/models/countriesNow/request"
	"country-rest-api/external/models/countriesNow/response"
	"country-rest-api/util"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RequestInfo(country string, r *http.Request) []string {

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	cityResponse := response.CityResponse{}

	jsonBody := util.EncodeJSONBody(request.PostRequestBody{Name: country})

	fmt.Println(request.PostRequestBody{Name: country})

	fmt.Println(jsonBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, constants.CountriesNowAPI+"countries/cities", jsonBody)
	if err != nil {
		log.Printf(constants.ErrorCreateRequest, err)
		return make([]string, 0)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Printf(constants.ErrorResponse, err2)
		return make([]string, 0)
	}

	fmt.Println(resp.Body)

	defer func(Body io.ReadCloser) {
		err3 := Body.Close()
		if err3 != nil {
			log.Printf(constants.ErrorCloseBody, err3)
		}
	}(resp.Body)

	body, err3 := io.ReadAll(resp.Body)
	if err3 != nil {
		log.Printf(constants.ErrorReadBody, err3)
		return make([]string, 0)
	}

	util.DecodeJSONBody(body, &cityResponse)
	fmt.Println(cityResponse)

	return cityResponse.Data
}
