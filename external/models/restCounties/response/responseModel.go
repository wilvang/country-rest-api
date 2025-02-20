package response

// Country represents information about a country.
type Country struct {
	Flags      Flags             `json:"flags"`
	Name       Name              `json:"name"`
	Capital    []string          `json:"capital"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Population int               `json:"population"`
	Continents []string          `json:"continents"`
}

// Flags represents the flag information of a country.
type Flags struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
	Alt string `json:"alt"`
}

// Name represents the name information of a country.
type Name struct {
	Common     string                `json:"common"`
	Official   string                `json:"official"`
	NativeName map[string]NativeName `json:"nativeName"`
}

// NativeName represents the native name information of a country.
type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}
