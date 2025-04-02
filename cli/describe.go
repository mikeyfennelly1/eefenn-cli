package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
)

func Describe(commandName string) error {
	c, err := core.GetCore()
	if err != nil {
		return err
	}

	command, err := c.GetCommandByName(commandName)
	if err != nil {
		return err
	}
	//bold := \033[1m
	//reset := \033[0m

	fmt.Printf("\033[1m%s:\033[0m", command.Name)

	fmt.Printf("\n\n")

	fmt.Printf("\033[1m%s\033[0m", "DESCRIPTION:")

	fmt.Printf("\n")
	fmt.Printf("%s\n", command.Description)

	fmt.Printf("\n\033[1m%s\033[0m\n", "COMMAND ARGUMENTS:")

	if len(command.Args) > 0 {
		for _, arg := range command.Args {
			fmt.Printf("\033[1m%s\033[0m%s", arg.Name, arg.Type)
			fmt.Printf("\n\n")
			fmt.Printf("%s", arg.Description)
		}
	} else {
		fmt.Printf("This command has no arguments.\n")
	}
	return nil
}
