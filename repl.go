package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/srinandan33/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			ip := scanner.Text()
			cleaned := cleanInput(ip)
			if len(ip) == 0 {
				continue
			}
			command, args := cleaned[0], cleaned[1:]
			c, ok := cfg.commands[command]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := c.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func buildCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Fetch next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Fetch previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area to see pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
	}
}

type cursor struct {
	previous *string
	next     *string
}

type config struct {
	commands       map[string]cliCommand
	locationCursor cursor
	pokeApiClient  *pokeapi.PokeAPIClient
	pokedex        map[string]pokeapi.Pokemon
}
