package cmd

import (
	"fmt"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Pokedex(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Your Pokédex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
