package handler

import (
	"country-rest-api/constants"
	"log"
	"net/http"
	"os"
)

// StartPage serves as the default handler for the root path.
// It informs the client that no functionality is provided at the root level
// and provides links to other available paths.
func StartPage(w http.ResponseWriter, r *http.Request) {

	// Read the HTML file
	htmlFile, err := os.ReadFile("frontend/index.html")
	if err != nil {
		log.Printf(constants.ErrorReadingHTML+" %v", err)
		http.Error(w, constants.ErrorReadingHTML, http.StatusInternalServerError)
		return
	}

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML file content to the response
	_, err2 := w.Write(htmlFile)
	if err2 != nil {
		log.Printf(constants.ErrorWritingHTML+" %v", err2)
		http.Error(w, constants.ErrorWritingHTML, http.StatusInternalServerError)
	}
}

// RedirectHandler handles HTTP requests and redirects them to a new URL.
// w: The HTTP response writer.
// r: The HTTP request.
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the new URL
	http.Redirect(w, r, constants.DefaultPath, http.StatusFound)
}
