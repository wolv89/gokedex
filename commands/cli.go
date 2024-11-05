package commands

import (
	"fmt"
	"strings"
)

type Command struct {
	name        string
	description string
	Callback    func() error
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
