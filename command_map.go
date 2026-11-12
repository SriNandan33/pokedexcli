package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	if cfg.locationCursor.next == nil || *cfg.locationCursor.next == "" {
		return errors.New("You are on last page!")
	}

	url := *cfg.locationCursor.next
	locationResponse, err := cfg.pokeApiClient.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.locationCursor.next = &locationResponse.Next
	cfg.locationCursor.previous = &locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.locationCursor.previous == nil || *cfg.locationCursor.previous == "" {
		return errors.New("You are on first page")
	}

	url := *cfg.locationCursor.previous
	locationResponse, err := cfg.pokeApiClient.FetchLocations(url)
	if err != nil {
		return err
	}

	cfg.locationCursor.next = &locationResponse.Next
	cfg.locationCursor.previous = &locationResponse.Previous

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}