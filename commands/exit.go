package commands

import (
	"fmt"
	"os"
)

func Exit(config *Config) error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
