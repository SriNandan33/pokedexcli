package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			command := cleaned[0]
			c, ok := cfg.commands[command]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			err := c.callback(cfg)
			if err != nil {
				fmt.Println("Error executing the command", err)
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
	callback    func(cfg *config) error
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
	}
}

type config struct {
	commands map[string]cliCommand
}
