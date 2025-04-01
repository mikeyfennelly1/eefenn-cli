package utils

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/eefenn/eefenn-cli/core/config"
)

// GetCommand
//
// # Gets a command by command name
//
// Returns
// - ptr to index of the commmand
// - ptr to the Subcommand structure for the command
// - error status
func GetCommand(commandName string) (*int, *cmd_config.Command, error) {
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return nil, nil, err
	}

	for index, command := range currentConfig.Commands {
		if command.Name == commandName {
			return &index, &command, nil
		}
	}

	return nil, nil, fmt.Errorf("Could not find the command %s\n", commandName)
}
