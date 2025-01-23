package commands

import (
	"fmt"

	"github.com/wolv89/gokedex/pokeapi"
)

func Explore(config *Config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("Please enter the name of an area to explore!")
	}

	pokemon, err := pokeapi.GetPokemon(args[0])
	if err != nil {
		return err
	}

	for _, pkm := range pokemon {
		fmt.Println(pkm.Name)
	}

	return nil

}
