package commands

import (
	"fmt"
)

func Pokedex(config *Config, args []string) error {

	if len(*config.Dex) == 0 {
		return fmt.Errorf("You haven't caught any Pokemon!")
	}

	fmt.Println("Your Pokedex:")

	for name := range *config.Dex {
		fmt.Printf(" - %s\n", name)
	}

	return nil

}
