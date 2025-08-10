package pokeapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	fullURL := BaseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var areaMap RespShallowLocations

	client := NewClient(5 * time.Second)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to make request: %w", err)
	}

	resp, err := client.HTTP.Do(req)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&areaMap); err != nil {
		return RespShallowLocations{}, err
	}

	return areaMap, nil
}
