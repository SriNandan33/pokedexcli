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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			ip := scanner.Text()
			cleaned := cleanInput(ip)
			if len(ip) == 0 {
				continue
			}
			firstWord := cleaned[0]
			fmt.Printf("Your command was: %s\n", firstWord)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
	}
}
