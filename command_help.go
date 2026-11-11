package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	supportedCommands := buildCommandRegistry()
	for _, cmd := range supportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
