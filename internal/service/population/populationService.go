package population

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service/countriesNow"
	"country-rest-api/external/api/service/restCountries"
	"country-rest-api/models"
	"net/http"
	"strconv"
	"strings"
)

// RequestPopulationService retrieves population data for a given country code and applies a year limit if provided.
// param: The country code (e.g., "US").
// limit: A string representing the year range in the format "fromYear-toYear" (e.g., "2010-2021").
// r: The HTTP request object.
// Returns a Population struct containing the mean population and historical population data.
func RequestPopulationService(param string, limit string, r *http.Request) models.Population {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.CountryFilter

	// Send request to REST Countries API and retrieve country information
	nameResponse := restCountries.RequestInfo(url, r)
	// Send request to Countries Now API and retrieve city information
	history := countriesNow.RequestPopulation(nameResponse.Name.Common, r).Data.PopulationCounts

	// Apply year limit if provided
	if limit != "" {
		history = trimHistory(limit, history)
	}

	// Populate the Population struct with the retrieved data
	population := models.Population{
		Mean:    0,
		History: history,
	}

	// Calculate the mean population
	sum := 0
	for _, population := range history {
		sum += population.Count
	}

	// Populate the mean attribute if the history isn't empty
	if len(history) != 0 {
		population.Mean = sum / (len(history))
	} else {
		population.History = []models.PopulationData{}
	}

	return population
}

// trimHistory filters the population history based on the provided year range.
// limit: A string representing the year range in the format "fromYear-toYear" (e.g., "2010-2021").
// history: A slice of PopulationData structs representing the historical population data.
// Returns a filtered slice of PopulationData structs.
func trimHistory(limit string, history []models.PopulationData) []models.PopulationData {

	// Split the limit string into fromYear and toYear using the "-" delimiter
	years := strings.Split(limit, "-")
	if len(years) != 2 {
		// If the limit string does not contain exactly two parts, return the original history
		return history
	}

	// Convert the fromYear and toYear strings to integers
	fromYear, err1 := strconv.Atoi(years[0])
	toYear, err2 := strconv.Atoi(years[1])
	if err1 != nil || err2 != nil {
		// If there is an error converting either year, return the original history
		return history
	}

	// Initialize a slice to hold the filtered population data
	var filteredPopulations []models.PopulationData

	// Iterate through the history slice
	for _, population := range history {
		// Check if the population year is within the specified range
		if population.Year >= fromYear && population.Year <= toYear {
			// If it is, append the population data to the filteredPopulations slice
			filteredPopulations = append(filteredPopulations, population)
		}
	}

	return filteredPopulations
}
