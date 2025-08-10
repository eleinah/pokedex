package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := BaseURL + "/pokemon/" + pokemonName

	var pokemon Pokemon

	if val, ok := c.Cache.Get(fullURL); ok {
		pokemon := Pokemon{}
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.HTTP.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.Cache.Add(fullURL, data)
	return pokemon, nil
}
