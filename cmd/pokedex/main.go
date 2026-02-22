package main

import (
	"time"

	"github.com/danarrigo/pokedex/internal/pokeapi" 
	"github.com/danarrigo/pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(30 * time.Second)

	cfg := &repl.Config{
		PokeapiClient: pokeClient,
	}

	repl.StartRepl(cfg)
}
