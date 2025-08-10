package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eleinah/pokedex/cmd"
	"github.com/eleinah/pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func startRepl(cfg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokédex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())
		command := strings.ToLower(userInput[0])

		switch command {
		case "exit":
			cmd.Exit(cfg)
		case "help":
			cmd.Help(cfg)
		case "map":
			cmd.Mapf(cfg)
		case "mapb":
			cmd.Mapb(cfg)
		default:
			fmt.Printf("error: command '%v' not found, type 'help' to see usage\n", command)
		}
	}
}
