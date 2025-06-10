package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (Ctrl + D to stop)")

	pagination := paginationConfig{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	} 

	for {
		fmt.Print("Pokedex > ")	

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		command, exists := commandsMap[words[0]]
		if exists {
			if err := command.callback(&pagination); err != nil {
				fmt.Println(err)
			}
			continue
		}else {
			fmt.Println("Unknown command")
			continue
		}
		}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type paginationConfig struct {
	Next string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*paginationConfig) error
}

var commandsMap = make(map[string]cliCommand)

func init() {
	commandsMap["exit"] = cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	}
	commandsMap["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
	commandsMap["map"] = cliCommand{
		name: "map",
		description: "Displays the names of 20 location areas in the Pokemon world.",
		callback: commandMap,
	}
	commandsMap["mapb"] = cliCommand{
		name: "mapb",
		description: "Displays the names of 20 previous location areas in the Pokemon world.",
		callback: commandMapb,
	}
}