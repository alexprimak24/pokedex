package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(pagination *paginationConfig) error {
	res, err := http.Get(pagination.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	locations := Locations{}
	decodeErr := json.Unmarshal(body, &locations)
	if decodeErr != nil {
		return decodeErr
	}
	// Current request becomes previous
	pagination.Previous = pagination.Next
	// Next request will be user for the next call
	pagination.Next = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil;
}