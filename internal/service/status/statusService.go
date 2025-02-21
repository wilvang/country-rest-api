package status

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service"
	"country-rest-api/models"
	"time"

	"net/http"
)

func RequestStatusService(r *http.Request) models.Status {
	return models.Status{
		CountriesNow:  service.RequestStatus(constants.CountriesNowStatus, r),
		RestCountries: service.RequestStatus(constants.RESTCountriesStatus, r),
		Version:       constants.Version,
		Uptime:        int(time.Since(constants.StartTime).Seconds()),
	}
}
