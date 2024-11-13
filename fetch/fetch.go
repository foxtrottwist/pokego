package fetch

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	BASE_URL          = "https://pokeapi.co/api/v2/"
	LOCATION_AREA_URL = BASE_URL + "location-area"
)

func LocationAreas(url string) (locationAreas, error) {
	locationAreaUrl := LOCATION_AREA_URL

	if url != "" {
		locationAreaUrl = url
	}

	res, err := http.Get(locationAreaUrl)
	if err != nil {
		return locationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	var locations locationAreas
	if err := json.Unmarshal(data, &locations); err != nil {
		return locationAreas{}, err
	}

	return locations, nil
}
