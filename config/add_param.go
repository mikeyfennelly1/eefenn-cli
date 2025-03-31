package config

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd-config"
)

// AddParam
//
// add a parameter to a command
func AddParam(commandName string, p cmd_config.Parameter) error {
	// get the current config
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return err
	}

	// find the subcommand with name commandName and update
	// the config as such if command is found
	for subCommandIndex, sc := range currentConfig.Subcommands {
		if sc.Name == commandName {
			// append the parameter to the subcommand's parameter array
			currentConfig.Subcommands[subCommandIndex].Parameters = append(sc.Parameters, p)
		} else {
			continue
		}
		// update the config file
		err = currentConfig.Update()
		if err != nil {
			return err
		}
		return nil
	}

	// if this point is reached, the command wasn't found, return error
	return fmt.Errorf("Could not find command %s\n", commandName)
}
