package main

import (
	"time"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func main() {
	timeoutInterval := 5 * time.Second
	cacheInterval := 5 * time.Minute
	client := pokeapi.NewClient(timeoutInterval, cacheInterval)

	cfg := &pokeapi.Config{
		PokeAPIClient: client,
	}

	startRepl(cfg)
}
