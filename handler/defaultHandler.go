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
		log.Printf("Error writing HTML file to response: %v", err)
		http.Error(w, "Error reading HTML file", http.StatusInternalServerError)
		return
	}

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML file content to the response
	_, err2 := w.Write(htmlFile)
	if err2 != nil {
		log.Printf("Error writing HTML file to response: %v", err2)
		http.Error(w, "Error writing HTML file to response", http.StatusInternalServerError)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect to the new URL
	http.Redirect(w, r, constants.DefaultPath, http.StatusFound)
}
