package repl

import (
	"fmt"
	"os"
	"sort"
	"math/rand"
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
		"explore":{
			name:		"explore",
			description:"Displays list of pokemon in a certain area",
			callback:	commandExplore,
		},
		"catch":{
			name:		"catch",
			description:"Catches a specific pokemon",
			callback:	commandCatch,
		},
	}
}

func commandExit(cfg *Config,strs []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config,strs []string) error {
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
	fmt.Println()
	return nil
}

func commandMap(cfg *Config,strs []string) error {
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

func commandMapb(cfg *Config,strs[]string) error {
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

func commandExplore(cfg *Config,strs[]string)error{
	resp,err:=cfg.PokeapiClient.FetchSpecificLocationInfo(strs[1])
	if err!=nil{
		return err
	}
	fmt.Printf("Exploring %s... ",strs[1])
	fmt.Printf("Found Pokemon: \n")
	for _, encounter := range resp.PokemonEncounters{
		fmt.Printf(" - %v\n",encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *Config,strs[]string)error{
	res,err:=cfg.PokeapiClient.GetPokemonInfo(strs[1])
	if err!=nil{
		return err
	}
	base_xp:=res.BaseExperience
	fmt.Printf("Throwing a Pokeball at %s...",res.Name)
	chance:=rand.Intn(base_xp)
	threshold := 50
	if chance < threshold {
	    fmt.Printf("%s was caught!\n",res.Name)
	    cfg.PokeDex[res.Name] = res
	} else {
	    fmt.Printf("%s escaped!\n",res.Name)
	}
	
	return nil
}
