package core

import (
	"github.com/eefenn/eefenn-cli/core/command_dir"
	"github.com/eefenn/eefenn-cli/core/config"
	cmd_config "github.com/eefenn/eefenn-cli/yaml"
)

type CoreInterface interface {
	// CommitCommand
	//
	// Add/'commit' a command to core.
	CommitCommand(command cmd_config.Command)

	// GetCommandByName
	//
	// Get a Command object for a command, using the name of the command as
	// a parameter.
	GetCommandByName(commandName string) (cmd_config.Command, error)

	// GetALlCommands
	//
	// Get all commands in the current core state.
	GetALlCommands() []cmd_config.Command

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
func (c *Core) Commit(command cmd_config.Command) {

}

func (c *Core) GetCommands() []cmd_config.Command {
	return nil
}
