package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

func FetchLocations(url string) (locationResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return locationResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse{}, err
	}

	resp.Body.Close()
	if resp.StatusCode != 200 {
		return locationResponse{}, errors.New("Error fetching locations")
	}

	lr := locationResponse{}
	err = json.Unmarshal(body, &lr)
	if err != nil {
		return locationResponse{}, err
	}

	return lr, nil
}
