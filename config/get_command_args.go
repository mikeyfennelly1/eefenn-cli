package config

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd-config"
)

// GetCommandArgs
//
// Get the arguments of a command by command name
func GetCommandArgs(commandName string) ([]cmd_config.Arg, error) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	for _, sc := range currentConfig.Commands {
		if sc.Name == commandName {
			return sc.Args, nil
		}
	}

	return nil, fmt.Errorf("Could not find parameters for the command: '%s'\n", commandName)
}
