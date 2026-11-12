package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/srinandan33/pokedexcli/internal/pokecache"
)

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationResponse struct {
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

func (pc *PokeAPIClient) FetchLocations(url string) (locationResponse, error) {
	var body []byte
	fromCache, ok := pc.cache.Get(url)
	if ok {
		body = fromCache
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return locationResponse{}, err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return locationResponse{}, err
		}

		resp.Body.Close()
		if resp.StatusCode != 200 {
			return locationResponse{}, errors.New("Error fetching locations")
		}
		pc.cache.Add(url, body)
	}

	lr := locationResponse{}
	err := json.Unmarshal(body, &lr)
	if err != nil {
		return locationResponse{}, err
	}

	return lr, nil
}
