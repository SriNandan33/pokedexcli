package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/srinandan33/pokedexcli/internal/pokecache"
)

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ListLocationsResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []location `json:"results"`
}

type PokeAPIClient struct {
	cache *pokecache.Cache
}

func CreatePokeApiClient() *PokeAPIClient {
	return &PokeAPIClient{
		cache: pokecache.NewCache(5 * time.Minute),
	}
}

func (pc *PokeAPIClient) FetchLocations(url string) (ListLocationsResponse, error) {
	var body []byte
	fromCache, ok := pc.cache.Get(url)
	if ok {
		body = fromCache
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return ListLocationsResponse{}, err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return ListLocationsResponse{}, err
		}

		resp.Body.Close()
		if resp.StatusCode != 200 {
			return ListLocationsResponse{}, errors.New("Error fetching locations")
		}
		pc.cache.Add(url, body)
	}

	lr := ListLocationsResponse{}
	err := json.Unmarshal(body, &lr)
	if err != nil {
		return ListLocationsResponse{}, err
	}

	return lr, nil
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationAreaResponse struct {
	// Other properties are ignored as we don't use them
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func (pc *PokeAPIClient) FetchALocation(name string) (LocationAreaResponse, error) {
	var body []byte
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	fromCache, ok := pc.cache.Get(url)
	if ok {
		fmt.Println("Fetching from cache")
		body = fromCache
	} else {
		fmt.Println("Fetched from network")
		resp, err := http.Get(url)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		resp.Body.Close()
		if resp.StatusCode != 200 {
			return LocationAreaResponse{}, errors.New("Error fetching locations")
		}
		pc.cache.Add(url, body)
	}

	lr := LocationAreaResponse{}
	err := json.Unmarshal(body, &lr)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return lr, nil
}
