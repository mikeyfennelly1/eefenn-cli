package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/config"
)

func Describe(commandName string) error {
	description, err := config.GetCommandDescription(commandName)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *description)
	return nil
}
