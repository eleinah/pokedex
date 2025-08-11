package cmd

import "github.com/eleinah/pokedex/internal/pokeapi"

type cliCmd struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
}

func GetCmds() map[string]cliCmd {
	return map[string]cliCmd{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokédex",
			Callback:    Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    Help,
		},
		"clear": {
			Name:        "clear",
			Description: "Clear the display",
			Callback:    Clear,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    Mapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    Mapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore a location",
			Callback:    Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    Catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught pokemon",
			Callback:    Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "View your Pokédex",
			Callback:    Pokedex,
		},
	}
}
