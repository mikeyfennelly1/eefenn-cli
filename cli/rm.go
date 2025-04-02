package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
)

// RM
//
// Remove a command by name
func RM(commandName string) error {
	if !core.CMDExists(commandName) {
		return fmt.Errorf("Command '%s' does not exist.", commandName)
	}

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
