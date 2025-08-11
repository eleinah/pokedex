package cmd

import (
	"fmt"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Clear(cfg *pokeapi.Config, args ...string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}
