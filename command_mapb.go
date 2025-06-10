package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
)

func commandMapb(pagination *paginationConfig) error {
	if pagination.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(pagination.Previous)
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

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil;
}