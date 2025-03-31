package config

import (
	"fmt"
)

// RmParam
//
// remove a parameter from a command
func RmParam(commandName string, parameterName string) error {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return err
	}

	// find the command with name=commandName
	for subcommandIndex, sc := range currentConfig.Subcommands {
		if sc.Name == commandName {
			// find the parameter of the command if there is one,
			// where parameter.name == parameterName
			for index, p := range sc.Parameters {
				if p.Name == parameterName {
					// update the config
					currentConfig.Subcommands[subcommandIndex].Parameters = append(sc.Parameters[:index], sc.Parameters[index+1:]...)
				}
				currentConfig.Update()
				return nil
			}
			return fmt.Errorf("Command '%s' has no parameter '%s'\n", commandName, parameterName)
		}
	}

	return fmt.Errorf("Could not find command with name '%s'\n", commandName)
}
