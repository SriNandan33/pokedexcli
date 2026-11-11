package main

import (
	"errors"
	"fmt"

	"github.com/srinandan33/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	if cfg.locations.next == nil {
		return errors.New("You are on last page!")
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

func commandMapb(cfg *config) error {
	if cfg.locations.previous == nil {
		return errors.New("You are on first page")
	}

	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	if *cfg.locations.previous != "" {
		url = *cfg.locations.previous
	}

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
