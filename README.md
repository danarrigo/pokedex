# Pokedex CLI

A command-line interface (CLI) Pokedex application written in Go. It uses a REPL (Read-Eval-Print Loop) to allow users to interact with the PokeAPI, explore locations, and catch Pokemon.

## Features and Commands

The application provides several interactive commands inside the REPL:

* **help**: Displays a help message with a list of available commands.
* **exit**: Exits the Pokedex application.
* **map**: Displays the next 20 locations of the Pokemon world by paginating through the API.
* **mapb**: Displays the previous 20 locations.
* **explore <location_name>**: Displays a list of Pokemon found in a specific location area.
* **catch <pokemon_name>**: Attempts to catch a specific Pokemon. The catch rate is determined by a random chance calculated against the Pokemon's base experience.
* **inspect <pokemon_name>**: Inspects a Pokemon that you have successfully caught, displaying its height, weight, and types.
* **pokedex**: Shows a list of all the Pokemon that have been caught and added to your personal Pokedex map.

## Architecture

* **REPL**: Runs a continuous loop taking user input from the standard input (`stdin`), parsing it into clean commands, and executing the corresponding callbacks.
* **PokeAPI Client**: Handles HTTP requests to the external PokeAPI with a configurable timeout of 30 seconds.
* **Caching Mechanism**: The client includes an internal cache configured to clear data on 5-minute intervals, which helps avoid redundant network calls during your session.
* **Local State**: Maintains a local Pokedex state in memory using a map (`map[string]pokeapi.PokemonInfo`) to keep track of the Pokemon you have caught.

## Getting Started

To run the application, navigate to the root directory of the project and execute the main package:

```bash
go run ./cmd/pokedex
