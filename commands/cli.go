package commands

import (
	"fmt"
	"strings"

	"github.com/wolv89/gokedex/pokeapi"
)

type Command struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

type Config struct {
	Next     string
	Previous string
	Dex      *map[string]pokeapi.CaughtPokemon
}

func GetCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"map": {
			name:        "map",
			description: "List the next 20 locations from the PokeAPI",
			Callback:    MapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations from the PokeAPI",
			Callback:    MapBack,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "List all of the Pokemon in a given location area",
			Callback:    Explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Throw a Pokeball at a Pokemon and try to catch it",
			Callback:    Catch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "List the stats of a Pokemon in your Pokedex",
			Callback:    Inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of the Pokemon in your Pokedex",
			Callback:    Pokedex,
		},
	}
}

func ListCommands() {
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
}

func CleanInput(raw string) []string {
	text := strings.ToLower(raw)
	return strings.Fields(text)
}
