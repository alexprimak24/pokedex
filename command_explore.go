package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	fmt.Println("Exploring pastoria-city-area...")
	locationResp, err := cfg.pokeapiClient.ListLocPokemons(params)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")	
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println("- ", pokemon.Pokemon.Name)
	}
	return nil

}

