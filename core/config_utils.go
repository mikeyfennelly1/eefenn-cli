package core

import "github.com/eefenn/eefenn-cli/cmd"

// CMDExistsInConfig
//
// Check if the config file contains an entry for the command 'cmd'
func CMDExistsInConfig(cmd cmd.Command) (int, bool) {
	currentConfig, _ := getCurrentConfig()

	for index, existingCommand := range currentConfig.Commands {
		if existingCommand.Name == cmd.Name {
			return index, true
		}
	}

	return -1, false
}
