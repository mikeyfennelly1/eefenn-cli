// Subcommand.go
//
// asc (add Subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package commands

import (
	command_dir2 "github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/google/uuid"
)

// AddSubCommand
//
// Add a Subcommand, and it's script to the user's CLI
func AddSubCommand(sc *subcommand.Subcommand) error {
	// create directory structure
	err := command_dir2.CreateSubcommandDirTree(sc.Hash)
	if err != nil {
		return err
	}

	// copy the shell script
	err = command_dir2.CopyShellScript(sc.Script, sc.Hash)
	if err != nil {
		return err
	}

	// update the eefenn-cli.config.json to contain the command info

	return nil
}

// RemoveSubcommand
//
// Remove a subcommand's directories by command hash
func RemoveSubcommand(commandHash string, commandName string) error {
	err := command_dir2.RemoveCommandDirectoryRecursively(commandHash)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// CreateSubCommand
//
// Create a Subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, description string) subcommand.Subcommand {
	UUID := uuid.New().String()
	subCommand := subcommand.Subcommand{
		Name:        name,
		Hash:        UUID,
		Script:      sourceScriptName,
		Description: description,
	}
	return subCommand
}
