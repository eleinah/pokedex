package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	fullURL := BaseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var areaMap RespShallowLocations

	if val, ok := c.Cache.Get(fullURL); ok {
		locations := RespShallowLocations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locations, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to make request: %w", err)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if err := json.Unmarshal(data, &areaMap); err != nil {
		return RespShallowLocations{}, err
	}

	c.Cache.Add(fullURL, data)
	return areaMap, nil
}
