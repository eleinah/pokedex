package cmd

import (
	"errors"
	"fmt"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Mapf(cfg *pokeapi.Config) error {
	locations, err := cfg.PokeAPIClient.ListLocations(cfg.NextLocationURL)
	if err != nil {
		return err
	}

	cfg.NextLocationURL = locations.Next
	cfg.PreviousLocationURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func Mapb(cfg *pokeapi.Config) error {
	if cfg.PreviousLocationURL == nil {
		err := errors.New("you're on the first page")
		fmt.Printf("error: %v\n", err)
		return err
	}

	locations, err := cfg.PokeAPIClient.ListLocations(cfg.PreviousLocationURL)
	if err != nil {
		return err
	}

	cfg.NextLocationURL = locations.Next
	cfg.PreviousLocationURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
