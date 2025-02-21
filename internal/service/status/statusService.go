package status

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service"
	"country-rest-api/models"
	"net/http"
	"time"
)

// RequestStatusService retrieves the status of external services and returns a Status struct.
// It checks the status of the Countries Now and REST Countries services, and includes the
// current version and uptime of the application.
func RequestStatusService(r *http.Request) models.Status {
	return models.Status{
		CountriesNow:  service.RequestStatus(constants.CountriesNowStatus, r),
		RestCountries: service.RequestStatus(constants.RESTCountriesStatus, r),
		Version:       constants.Version,
		Uptime:        int(time.Since(constants.StartTime).Seconds()),
	}
}
