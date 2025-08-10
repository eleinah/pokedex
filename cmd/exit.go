package cmd

import (
	"fmt"
	"os"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Exit(cfg *pokeapi.Config) error {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
	return nil
}
