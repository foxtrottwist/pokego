package client

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/foxtrottwist/pokego/cache"
)

type Client struct {
	httpClient http.Client
	cache.Cache
}

func New(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache: cache.New(5 * time.Second),
	}
}

func (c *Client) LocationAreas(url *string) (locationAreas, error) {
	locationAreaUrl := LOCATION_AREA_URL
	if url != nil {
		locationAreaUrl = *url
	}

	data, ok := c.Cache.Get(locationAreaUrl)

	if !ok {
		res, err := c.httpClient.Get(locationAreaUrl)
		if err != nil {
			return locationAreas{}, err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return locationAreas{}, err
		}

		data = body
		c.Cache.Add(locationAreaUrl, body)
	}

	var locations locationAreas
	if err := json.Unmarshal(data, &locations); err != nil {
		return locationAreas{}, err
	}

	return locations, nil
}
