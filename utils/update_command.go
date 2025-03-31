package utils

import (
	"github.com/eefenn/eefenn-cli/config"
	"github.com/eefenn/eefenn-cli/subcommand"
)

func ReplaceCommand(commandToReplace string, newCommand subcommand.Subcommand) error {
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
	currentConfig.Subcommands[*replaceTargetIndex] = newCommand

	// update the command
	err = currentConfig.Update()
	if err != nil {
		return err
	}

	return nil
}
