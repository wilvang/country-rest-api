package models

// Population represents the mean population and its historical data.
type Population struct {
	Mean    int              `json:"mean"`
	History []PopulationData `json:"values"`
}

// PopulationData represents population data for a specific year.
type PopulationData struct {
	Year  int `json:"year"`
	Count int `json:"values"`
}
