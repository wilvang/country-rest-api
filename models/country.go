package models

// Country represents information about a country.
type Country struct {
	// Country is the name of the country.
	Country string `json:"name"`

	// Continents is a list of continents the country is part of.
	Continents []string `json:"continents"`

	// Population is the population of the country.
	Population int `json:"population"`

	// Languages is a map of language codes to language names spoken in the country.
	Languages map[string]string `json:"languages"`

	// Borders is a list of countries that share a border with this country.
	Borders []string `json:"borders"`

	// Flag is the URL or path to the country's flag image.
	Flag string `json:"flag"`

	// Capital is the name of the capital city of the country.
	Capital string `json:"capital"`

	// Cities is a list of all cities in the country.
	Cities []string `json:"cities"`
}
