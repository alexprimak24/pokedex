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
			if err := command.callback(); err != nil {
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
}