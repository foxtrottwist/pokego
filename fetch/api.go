package fetch

const (
	BASE_URL          = "https://pokeapi.co/api/v2"
	LOCATION_AREA_URL = BASE_URL + "/location-area"
)

type locationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
