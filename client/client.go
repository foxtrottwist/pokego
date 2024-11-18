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
	httpClient http.Client
	cache      cache.Cache
}

func New(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
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

func (c *Client) LocationArea(name string) (LocationAreasResp, error) {
	locationAreaUrl := LOCATION_AREA_URL + "/" + name
	data, ok := c.cache.Get(locationAreaUrl)

	if !ok {
		res, err := c.httpClient.Get(locationAreaUrl)
		if err != nil {
			return LocationAreasResp{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAreasResp{}, errors.New("unable to explore area, check area name")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasResp{}, err
		}

		data = body
		c.cache.Add(locationAreaUrl, data)
	}

	var location LocationAreasResp
	if err := json.Unmarshal(data, &location); err != nil {
		return LocationAreasResp{}, err
	}

	return location, nil
}

func (c *Client) LocationAreas(url *string) (LocationAreasTruncResp, error) {
	locationAreaUrl := LOCATION_AREA_URL + LOCATION_AREA_DEFAULT_PARAMS
	if url != nil {
		locationAreaUrl = *url
	}

	data, ok := c.cache.Get(locationAreaUrl)

	if !ok {
		res, err := c.httpClient.Get(locationAreaUrl)
		if err != nil {
			return LocationAreasTruncResp{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationAreasTruncResp{}, errors.New("unable to display location areas")
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreasTruncResp{}, err
		}

		data = body
		c.cache.Add(locationAreaUrl, body)
	}

	var locations LocationAreasTruncResp
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreasTruncResp{}, err
	}

	return locations, nil
}
