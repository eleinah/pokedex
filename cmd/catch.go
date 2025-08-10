package cmd

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/eleinah/pokedex/internal/pokeapi"
)

func Catch(cfg *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeAPIClient.GetPokemon(name)
	if err != nil {
		return err
	}

	result := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if result > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}
