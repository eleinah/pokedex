package location

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eleinah/pokedex/internal/pokeapi"
	"github.com/eleinah/pokedex/internal/pokeapi/types"
)

func ListLocation() ([]types.RespShallowLocations, error) {
	var fullURL string = pokeapi.BaseURL + "/location-area"

	var areaMap []types.RespShallowLocations

	client := pokeapi.NewClient(100)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return []types.RespShallowLocations{}, fmt.Errorf("failed to make request: %w", err)
	}

	resp, err := client.HTTP.Do(req)
	if err != nil {
		return []types.RespShallowLocations{}, fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&areaMap); err != nil {
		return []types.RespShallowLocations{}, err
	}

	return areaMap, nil
}
