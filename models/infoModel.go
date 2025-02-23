package models

// Info represents information about a country.
type Info struct {
	Country    string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capitals   []string          `json:"capital"`
	Cities     []string          `json:"cities"`
}
