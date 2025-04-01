// Subcommand.go
//
// asc (add Subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
	"github.com/eefenn/eefenn-cli/core/config"
)

// RemoveSubcommand
//
// Remove a subcommand's directories by command hash
func RemoveSubcommand(commandName string) error {
	// get the current eefenn-cli.config.json
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	// remove the command directory for this command
	err = command_dir.RemoveCommandDirectoryRecursively(commandName)
	if err != nil {
		return err
	}

	// remove the entry for the command from the config file
	err = currentConfig.RemoveCommandByName(commandName)
	if err != nil {
		return err
	}

	// update the eefenn-cli.config.json
	currentConfig.update()

	fmt.Printf("Removed command: '%s'\n", commandName)

	return nil
}
