package cli

import (
	"github.com/eefenn/eefenn-cli/core"
)

// RM
//
// Remove a command by name
func RM(commandName string) error {
	err := core.RemoveCommandByName(commandName)
	if err != nil {
		return err
	}

	return nil
}
