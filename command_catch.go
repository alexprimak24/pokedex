package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params []string) error {
	pokemonName := params[0]
	// checking whether it is already caugth
	_, alreadyCaugth := usersPokedex[pokemonName]
	if alreadyCaugth {
		fmt.Printf("%s has been already caugth!\n", pokemonName)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonInfo, err := cfg.pokeapiClient.PokemonInfo(params)
	if err != nil {
		return err
	}

	experience := int(pokemonInfo.BaseExperience / 50)
	randomNum := rand.Intn(experience + 1)

	if randomNum == experience {
		usersPokedex[pokemonName] = pokemonInfo
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

