package cmd

import (
	"fmt"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Help(cfg *pokeapi.Config) error {
	fmt.Printf(`
Welcome to the Pokédex!
Usage:


`)
	for _, cmd := range getCmds() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
