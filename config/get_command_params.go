package config

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/subcommand"
)

func GetCommandParams(commandName string) ([]subcommand.Parameter, error) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	for _, sc := range currentConfig.Subcommands {
		if sc.Name == commandName {
			return sc.Parameters, nil
		}
	}

	return nil, fmt.Errorf("Could not find parameters for the command: '%s'\n", commandName)
}
