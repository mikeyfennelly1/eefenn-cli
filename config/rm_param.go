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

	for subcommandIndex, sc := range currentConfig.Subcommands {
		if sc.Name == commandName {
			for index, p := range sc.Parameters {
				if p.Name == parameterName {
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
