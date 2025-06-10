package main

import "fmt"

func commandHelp(pagination *paginationConfig) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n")
	fmt.Println()

	for _, command := range commandsMap {
		fmt.Println(command.name + ": " + command.description)
	}
	println()
	return nil
}