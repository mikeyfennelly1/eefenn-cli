package cli

import (
	"github.com/eefenn/eefenn-cli/core"
)

// RM
//
// Remove a command by name
func RM(commandName string) error {
	currentCore, err := core.GetCore()
	if err != nil {
		return err
	}

	err = currentCore.RemoveCommandByName(commandName)
	if err != nil {
		return err
	}

	return nil
}
