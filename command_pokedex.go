package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedex := cfg.pokedex
	if len(pokedex) == 0 {
		return errors.New("Your pokedex is empty! Go, Catch some")
	}
	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
