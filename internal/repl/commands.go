package repl

import (
	"fmt"
	"os"
	"sort"
)

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays locations of the pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	cliCommands := getCliCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	keys := make([]string, 0, len(cliCommands))
	for k := range cliCommands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		command := cliCommands[k]
		fmt.Printf("\n%s: %s", command.name, command.description)
	}
	fmt.Println("\n")
	return nil
}

func commandMap(cfg *Config) error {
	response, err := cfg.PokeapiClient.FetchLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	for _, val := range response.Results {
		fmt.Printf("%s\n", val.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	response, err := cfg.PokeapiClient.FetchLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = response.Next
	cfg.Previous = response.Previous

	for _, val := range response.Results {
		fmt.Printf("%s\n", val.Name)
	}
	return nil
}
