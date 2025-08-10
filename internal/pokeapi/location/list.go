package location

import (
	"encoding/json"
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
		return []types.RespShallowLocations{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []types.RespShallowLocations{}, err
	}
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&areaMap); err != nil {
		return []types.RespShallowLocations{}, err
	}

	return areaMap, nil
}
