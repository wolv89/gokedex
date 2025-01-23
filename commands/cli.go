package commands

import (
	"fmt"
	"strings"
)

type Command struct {
	name        string
	description string
	Callback    func(*Config, []string) error
}

type Config struct {
	Next     string
	Previous string
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
			name:        "explore",
			description: "List all of the Pokemon in a location area",
			Callback:    Explore,
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
