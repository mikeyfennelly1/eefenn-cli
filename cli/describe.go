package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/utils"
)

func Describe(commandName string) error {
	_, command, err := utils.GetCommand(commandName)
	if err != nil {
		return err
	}

	// print the command name
	fmt.Printf("\u001B[1m%s\u001B[0m\n\n\n", commandName)

	// print long form description paragraph
	fmt.Printf("\u001B[1m%s\u001B[0m\n", "DESCRIPTION:")
	fmt.Printf("%s\n\n\n", command.Description)

	// print command parameters
	fmt.Printf("\u001B[1m%s\u001B[0m \n\n", "PARAMETERS:")
	for _, param := range command.Args {
		// print `<param_name>:<param_description>`
		fmt.Printf("\u001B[1m%s\u001B[0m: %s", param.Name, param.Description)
	}

	fmt.Printf("\u001B[1m%s\u001B[0m \n\n", "SOURCE:")
	source, err := utils.GetScriptContents(commandName)
	fmt.Printf("%s\n", string(source))

	return nil
}
