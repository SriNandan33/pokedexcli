package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Invalid usage, expecting a name argument")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeApiClient.FetchPokemon(name)
	if err != nil {
		return err
	}
	catchChance := 100 - pokemon.BaseExperience/10
	caught := rand.Intn(100) < catchChance
	if caught {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
