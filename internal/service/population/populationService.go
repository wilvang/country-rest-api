package population

import (
	"country-rest-api/constants"
	"country-rest-api/external/api/service/countriesNow"
	"country-rest-api/external/api/service/restCountries"
	"country-rest-api/models"
	"country-rest-api/util"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// RequestPopulationService retrieves population data for a given country code and applies a year limit if provided.
// param: The country code (e.g., "US").
// limit: A string representing the year range in the format "fromYear-toYear" (e.g., "2010-2021").
// r: The HTTP request object.
// Returns a Population struct containing the mean population and historical population data.
func RequestPopulationService(param string, limit string, r *http.Request) (models.Population, error) {
	url := constants.RESTCountriesAPI + "alpha/" + param + constants.CountryFilter

	// Send request to REST Countries API and retrieve country information
	nameResponse := restCountries.RequestInfo(url, r)
	if util.IsEmpty(nameResponse) {
		return models.Population{}, errors.New(constants.ErrorNotFound)
	}

	// Send request to Countries Now API and retrieve city information
	populationHistory := countriesNow.RequestPopulation(nameResponse.Name.Common, r).Data.PopulationCounts
	if util.IsEmpty(populationHistory) {
		// If the response is empty, try a request with the official name
		populationHistory = countriesNow.RequestPopulation(nameResponse.Name.Official, r).Data.PopulationCounts
	}

	// Apply year limit if provided
	if limit != "" {
		populationHistory = trimHistory(limit, populationHistory)
	}

	// Populate the Population struct with the retrieved data
	population := models.Population{
		Mean:    0,
		History: populationHistory,
	}

	// Calculate the mean population
	sum := 0
	for _, population := range populationHistory {
		sum += population.Count
	}

	// Populate the mean attribute if the populationHistory isn't empty
	if len(populationHistory) != 0 {
		population.Mean = sum / (len(populationHistory))
	} else {
		population.History = []models.PopulationData{}
	}

	return population, nil
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
