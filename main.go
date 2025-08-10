package main

import (
	"time"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func main() {
	clientTimeout := 5 * time.Second
	client := pokeapi.NewClient(clientTimeout)

	cfg := &pokeapi.Config{
		PokeAPIClient: client,
	}

	startRepl(cfg)
}
