package commands

import (
	"github.com/eefenn/eefenn-cli/utils"
)

func AddDescription(commandName string, newDescription string) error {
	_, command, err := utils.GetCommand(commandName)
	if err != nil {
		return err
	}

	command.Description = newDescription
	err = utils.ReplaceCommand(commandName, *command)
	if err != nil {
		return nil
	}

	return err
}
