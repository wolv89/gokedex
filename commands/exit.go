package commands

import (
	"fmt"
	"os"
)

func Exit() error {
	fmt.Println("Exiting...")
	os.Exit(0)
	return nil
}
