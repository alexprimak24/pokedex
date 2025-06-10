package main

import "fmt"

func commandInspect(cfg *config, params []string) error {
	pokemonInfo, err := cfg.pokeapiClient.PokemonInfo(params)
	if err != nil {
		return err
	}

	fmt.Printf("Heigth: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Print("Stats:\n")
	if len(pokemonInfo.Stats) == 0 {
	fmt.Print("It seems like there is no stats :(\n")
	} else {
		for _, stats := range pokemonInfo.Stats {
			fmt.Printf("	- %s: %d\n", stats.Stat.Name, stats.BaseStat)
		}
	}
	fmt.Print("Types:\n")
	if len(pokemonInfo.Types) == 0 {
	fmt.Print("It seems like there is no types :(\n")
	} else {
		for _, pokeType := range pokemonInfo.Types {
			fmt.Printf("	- %s: %d\n", pokeType.Type.Name, pokeType.Slot)
		}
	}

	return nil
}