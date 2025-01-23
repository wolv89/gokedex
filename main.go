package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wolv89/gokedex/commands"
)

func main() {

	COMMANDS := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)

	config := commands.Config{}

	for {

		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		words := commands.CleanInput(input)
		if len(words) == 0 {
			continue
		}

		cmd := words[0]
		var args []string

		if len(words) > 1 {
			args = words[1:]
		}

		if _, ok := COMMANDS[cmd]; ok {

			err := COMMANDS[cmd].Callback(&config, args)

			if err != nil {
				fmt.Println(err)
			}

		} else {
			fmt.Println("Unknown command!")
		}

	}

}
