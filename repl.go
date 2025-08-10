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
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cmd.GetCmds()[commandName]
		if exists {
			err := command.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("error: unknown command")
			continue
		}
	}
}
