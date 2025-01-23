package commands

import (
	"fmt"
	"math/rand/v2"

	"github.com/wolv89/gokedex/pokeapi"
)

func Catch(config *Config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("Catch who?!")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := pokeapi.CatchPokemon(name)
	if err != nil {
		return err
	}

	throw := rand.IntN(1_000)

	if throw < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)

	(*config.Dex)[name] = pokemon

	return nil

}
