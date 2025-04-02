// core.go
//
// Core is essentially the management system for commands, and their files.
//
// Author: Mikey Fennelly

package core

import (
	"fmt"
	cmd "github.com/eefenn/eefenn-cli/cmd"
)

func GetCore() (CoreInterface, error) {
	config, err := getCurrentConfig()
	if err != nil {
		return nil, err
	}

	var edt eefennCLIDirectoryTree

	current_core := Core{
		config:        config,
		directoryTree: edt,
	}

	return &current_core, nil
}

type CoreInterface interface {
	// Commit
	//
	// Add/'commit' a command to core.
	Commit(command cmd.Command) error

	// GetCommandByName
	//
	// Get a Command object for a command, using the name of the command as
	// a parameter.
	GetCommandByName(commandName string) (*cmd.Command, error)

	// GetALlCommands
	//
	// Get all commands in the current core state.
	GetALlCommands() ([]cmd.Command, error)

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
	config        config
	directoryTree eefennCLIDirectoryTree
}

func (c *Core) GetCommandByName(commandName string) (*cmd.Command, error) {
	for _, command := range c.config.Commands {
		if command.Name == commandName {
			return &command, nil
		}
	}

	return nil, fmt.Errorf("could not find command: %s", commandName)
}

// GetALlCommands
// Gets all commands in config. If there are no commands,
// will return an error
func (c *Core) GetALlCommands() ([]cmd.Command, error) {
	if len(c.config.Commands) == 0 {
		return nil, fmt.Errorf("there are no commands")
	}
	return c.config.Commands, nil
}

func (c *Core) EditCommand(commandName string) {
	//TODO implement me
	panic("implement me")
}

func (c *Core) RunCommand(commandName string) {
	//TODO implement me
	panic("implement me")
}

// Commit
//
// Add/'commit' a command to core.
func (c *Core) Commit(command cmd.Command) error {
	if c == nil {
		return fmt.Errorf("Core is not properly initialized\n")
	}

	if CmdExists(c, command.Name) {
		return fmt.Errorf("Command already exists\n")
	}

	var edt eefennCLIDirectoryTree

	// Add the command to the config file
	err := c.config.addCMD(command)
	if err != nil {
		return err
	}
	// Create the directory tree for the command
	err = c.directoryTree.CreateCMDDirTree(command)
	if err != nil {
		return err
	}
	// Copy the script for the command from the pwd to the script
	// in newly created directory tree.
	err = c.directoryTree.CopyScriptToCMDDir(command)
	if err != nil {
		return err
	}

	err = c.directoryTree.CopyDependenciesToDependenciesDir(command)
	if err != nil {
		return err
	}

	err = edt.CreateCMDDirTree(command)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) RemoveCommandByName(commandName string) error {
	if c == nil {
		return fmt.Errorf("Core is not properly initialized\n")
	}

	if CmdExists(c, commandName) {
		return fmt.Errorf("Command already exists\n")
	}

	fmt.Println("Got past CMDExists...")

	currentConfig, err := getCurrentConfig()
	if err != nil {
		return err
	}

	err = currentConfig.removeCommandByName(commandName)
	if err != nil {
		return err
	}

	var edt eefennCLIDirectoryTree

	err = edt.RemoveCommandDirectoryRecursively(commandName)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) GetCommands() []cmd.Command {
	return nil
}
