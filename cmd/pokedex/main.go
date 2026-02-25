package main

import (
	"time"

	"github.com/danarrigo/pokedex/internal/pokeapi" 
	"github.com/danarrigo/pokedex/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(30 * time.Second,5*time.Minute)

	cfg := &repl.Config{
		PokeapiClient: pokeClient,
		PokeDex: make(map[string]pokeapi.PokemonInfo),
	}

	repl.StartRepl(cfg)
}
