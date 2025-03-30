// Subcommand.go
//
// asc (add Subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package commands

import (
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
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

	// get the hash code for the command
	commandHash, err := currentConfig.GetCommandHash(commandName)
	if err != nil {
		return err
	}

	// remove the command directory for this command
	err = command_dir.RemoveCommandDirectoryRecursively(*commandHash)
	if err != nil {
		return err
	}

	// remove the entry for the command from the config file
	err = currentConfig.RemoveCommandByName(commandName)
	if err != nil {
		return err
	}

	// update the eefenn-cli.config.json
	currentConfig.Update()

	return nil
}
