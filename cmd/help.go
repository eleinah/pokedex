package cmd

import (
	"fmt"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Help(cfg *pokeapi.Config, args ...string) error {
	fmt.Printf(`
Welcome to the Pokédex!
Usage:


`)
	for _, cmd := range GetCmds() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
