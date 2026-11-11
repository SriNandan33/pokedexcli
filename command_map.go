package main

import (
	"fmt"

	"github.com/srinandan33/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.locations.next == nil {
		fmt.Println("You are on last page!")
		return nil
	}

	url := *cfg.locations.next
	locationResponse, err := pokeapi.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.locations.next = &locationResponse.Next
	cfg.locations.previous = &locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}
