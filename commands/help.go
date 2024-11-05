package commands

import (
	"fmt"
)

func Help(config *Config) error {

	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	ListCommands()

	fmt.Println()

	return nil

}
