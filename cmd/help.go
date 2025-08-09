package cmd

import (
	"fmt"
)

func Help() error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:


`)
	for _, cmd := range getCmds() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
