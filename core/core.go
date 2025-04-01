// core.go
//
// Core is essentially the management system for commands, and their files.
//
// Author: Mikey Fennelly

package core

import (
	cmd "github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core/command_dir"
	"github.com/eefenn/eefenn-cli/core/config"
)

func GetCore() (*Core, error) {
	config, err := config.GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	var edt command_dir.EefennCLIDirectoryTree

	current_core := Core{
		Config:        config,
		DirectoryTree: edt,
	}

	return &current_core, nil
}

type CoreInterface interface {
	// Commit
	//
	// Add/'commit' a command to core.
	Commit(command cmd.Command)

	// GetCommandByName
	//
	// Get a Command object for a command, using the name of the command as
	// a parameter.
	GetCommandByName(commandName string) (cmd.Command, error)

	// GetALlCommands
	//
	// Get all commands in the current core state.
	GetALlCommands() []cmd.Command

	// RemoveCommandByName
	//
	// Remove a command, specifying which command by name of the command.
	RemoveCommandByName(commandName string) error

	// EditCommand
	//
	// Edit a command, specifying which command by name of the command.
	EditCommand(commandName string)

	// RunCommand
	//
	// Run a command, specifying which command by name of the command.
	RunCommand(commandName string)
}

type Core struct {
	Config        config.Config
	DirectoryTree command_dir.EefennCLIDirectoryTree
}

// Commit
//
// Add/'commit' a command to core.
func (c *Core) Commit(command cmd.Command) error {
	var edt command_dir.EefennCLIDirectoryTree

	err := edt.CreateCMDDirTree(command)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) GetCommands() []cmd.Command {
	return nil
}
