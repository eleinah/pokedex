package main

import (
	"bufio"
	"fmt"
	"github.com/eleinah/pokedex/cmd"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cmdInput := cleanInput(scanner.Text())

		switch strings.ToLower(cmdInput[0]) {
		case "exit":
			cmd.Exit()
		case "help":
			cmd.Help()
		}
	}
}
