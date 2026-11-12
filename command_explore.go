package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Invalid usage, expecting a name argument")
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)
	encounters, err := cfg.pokeApiClient.FetchALocation(name)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
