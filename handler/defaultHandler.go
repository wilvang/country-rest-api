package handler

import (
	"country-rest-api/config"
	"fmt"
	"net/http"
)

// EmptyHandler serves as the default handler for the root path.
// EmptyHandler informs the client that no functionality is provided at the root level
// and provides links to other available paths.
func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. Please use paths" +
		" <a href=\"" + config.INFO_PATH + "\">" + config.INFO_PATH + "</a> or " +
		" or <a href=\"" + config.POPULATION_PATH + "\">" + config.POPULATION_PATH + "</a>" +
		" or <a href=\"" + config.STATUS_PATH + "\">" + config.STATUS_PATH + "</a>."

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
