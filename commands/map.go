package commands

import (
	"errors"
	"fmt"

	"github.com/wolv89/gokedex/pokeapi"
)

func MapForward(config *Config, args []string) error {

	from := ""
	if config.Next != "" {
		from = config.Next
	}

	locations, next, prev, err := pokeapi.GetLocations(from)
	if err != nil {
		return err
	}

	config.Next = next
	config.Previous = prev

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	return nil

}

func MapBack(config *Config, args []string) error {

	if config.Previous == "" {
		return errors.New("can't map back from first page")
	}

	locations, next, prev, err := pokeapi.GetLocations(config.Previous)
	if err != nil {
		return err
	}

	config.Next = next
	config.Previous = prev

	for _, loc := range locations {
		fmt.Println(loc.Name)
	}

	return nil

}
