package commands

import (
	"fmt"
)

func Inspect(config *Config, args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("Which Pokemon would you like to inspect?")
	}

	name := args[0]

	if _, ok := (*config.Dex)[name]; !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	poke := (*config.Dex)[name]

	fmt.Printf("Name: %s\n", poke.Name)
	fmt.Printf("Height: %d\n", poke.Height)
	fmt.Printf("Weight: %d\n", poke.Weight)

	fmt.Println("Stats:")

	for _, stat := range poke.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, tp := range poke.Types {
		fmt.Printf(" - %s\n", tp.Type.Name)
	}

	return nil

}
