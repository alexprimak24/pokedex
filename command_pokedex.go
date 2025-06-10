package main

import "fmt"

func commandPokedex(cfg *config, params []string) error {

	if len(usersPokedex) == 0 {
		fmt.Print("Ooops seems like you don't own any pokemons yet\n")
		fmt.Print("Use `catch <pokemon_name>` to broaden your collection :) \n")
		return nil
	}

	fmt.Print("Your Pokedex\n")
	for pokemonName := range usersPokedex {
		fmt.Printf("- %s\n", pokemonName)
	}
	return nil
}