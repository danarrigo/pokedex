package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

type locationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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

func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i, val := range words {
		words[i] = strings.ToLower(val)
	}
	return words
}

func startRepl(cfg *config) {
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
				err := command.callback(cfg)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
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

func fetchLocations(url string) (locationResponse, error) {
	var response locationResponse
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	response, err := fetchLocations(*cfg.Previous)
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

func commandMap(cfg *config) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	}

	response, err := fetchLocations(url)
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
