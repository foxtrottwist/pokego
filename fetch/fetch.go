package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) LocationAreas(url *string) (locationAreas, error) {
	locationAreaUrl := LOCATION_AREA_URL
	if url != nil {
		locationAreaUrl = *url
	}

	res, err := c.httpClient.Get(locationAreaUrl)
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
