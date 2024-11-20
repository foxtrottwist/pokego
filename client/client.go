package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/foxtrottwist/pokego/cache"
)

type Client struct {
	client http.Client
	cache  cache.Cache
}

func New(timeout, interval time.Duration) Client {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
		cache: cache.New(interval),
	}
}

func (c *Client) CleanCache() string {
	return c.cache.Clean()
}

func (c *Client) ListCache() []string {
	return c.cache.LS()
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := LOCATION_AREA_URL + "/" + name
	data, ok := c.cache.Get(url)

	if !ok {
		res, err := c.client.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationArea{}, errors.New("unable to explore area, check area name")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationArea{}, err
		}

		data = body
		c.cache.Add(url, data)
	}

	var location LocationArea
	if err := json.Unmarshal(data, &location); err != nil {
		return LocationArea{}, err
	}

	return location, nil
}

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreaTrunc, error) {
	url := LOCATION_AREA_URL
	if pageUrl != nil {
		url = *pageUrl
	}

	data, ok := c.cache.Get(url)

	if !ok {
		res, err := c.client.Get(url)
		if err != nil {
			return LocationAreaTrunc{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAreaTrunc{}, errors.New("unable to display location areas")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaTrunc{}, err
		}

		data = body
		c.cache.Add(url, body)
	}

	var locations LocationAreaTrunc
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreaTrunc{}, err
	}

	return locations, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := POKEMON_URL + "/" + name
	data, ok := c.cache.Get(url)

	if !ok {
		res, err := c.client.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return Pokemon{}, errors.New("unable to catch Pokemon, check Pokemon name")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}

		data = body
		c.cache.Add(url, data)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
