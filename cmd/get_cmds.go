package cmd

import "github.com/eleinah/pokedex/internal/pokeapi"

type cliCmd struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCmds() map[string]cliCmd {
	return map[string]cliCmd{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    Exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    Help,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    Mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    Mapb,
		},
	}
}
