package utils

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/eefenn/eefenn-cli/core/config"
)

func ReplaceCommand(commandToReplace string, newCmd cmd_config.Command) error {
	// get the index of the command that you want to replace
	replaceTargetIndex, _, err := GetCommand(commandToReplace)
	if err != nil {
		return err
	}

	// get the current config
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return nil
	}

	// replace the command
	currentConfig.Commands[*replaceTargetIndex] = newCmd

	// update the command
	err = currentConfig.Update()
	if err != nil {
		return err
	}

	return nil
}
