package cmd

import (
	"fmt"
	"os"
)

func Exit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
