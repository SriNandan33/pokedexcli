package main

import (
	"github.com/srinandan33/pokedexcli/internal/pokeapi"
)

func main() {
	supportedCommands := buildCommandRegistry()

	firstPageUrl := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	config := config{
		commands:       supportedCommands,
		locationCursor: cursor{previous: nil, next: &firstPageUrl},
		pokeApiClient:  pokeapi.CreatePokeApiClient(),
	}
	startRepl(&config)
}
