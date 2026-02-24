package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/danarrigo/pokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config,[]string) error
}

func StartRepl(cfg *Config) {
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); ok {
			cleanCommand := cleanInput(scanner.Text())
			if len(cleanCommand) == 0 {
				continue
			}
			if command, exists := cliCommands[cleanCommand[0]]; exists {
				err := command.callback(cfg,cleanCommand)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, val := range words {
		words[i] = strings.ToLower(val)
	}
	return words
}
