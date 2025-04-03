// command_files_controller.go
//
// For interacting with the directory tree that manages commands.
//
// Author: Mikey Fennelly <mikeyp.fennelly@gmail.com>

package core

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// CommandFilesController
//
// For interfacing with the directory tree starting from the root: /usr/lib/eefenn-cli
//
// This can be used to
//   - Create subdirectories that correspond to the conventional directory tree
//     structure for an eefenn command.
//   - Utilities for getting data relating to the conventional directory tree such as
//     script paths for command scripts, dependency directory paths.
//   - Methods for interacting with a command's files, such as creating, removing and
//     moving command subdirectories to other directories in the filesystem for editing.
type CommandFilesController interface {
	// MoveCommand
	// Given a command and the command files current root (i.e the directory
	// where the script for the command is), this function moves all files from
	// <currentRootAbsPath>/<filePath> to <newRootAbsPath>/<filePath>
	MoveCommand(cmd cmd.CommandInterface) error

	// RemoveCommandRecursively
	//
	// remove a command directory recursively.
	RemoveCommandRecursively(commandName string) error
}

type cmdFilesController struct{}

// MoveCommand
// Move all contents of a command (script, config.yaml and dependencies).
func (cf *cmdFilesController) MoveCommand(cmd cmd.CommandInterface) error {
	//TODO implement me
	panic("implement me")
}

// CopyDependenciesToDependenciesDir
//
// Move a shell script to its command's directory
func (cf *cmdFilesController) CopyDependenciesToDependenciesDir(command cmd.CommandInterface) error {

	return nil
}

// RemoveCommandDirectoryRecursively
//
// remove a command directory recursively by command hash
func (cf *cmdFilesController) RemoveCommandRecursively(commandName string) error {
	dirname := getCommandDirAbsPath(commandName)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}

// createCMDDir
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func (cf *cmdFilesController) createCMDDir(command cmd.Command) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must have root permissions to perform changes to CLI core\n")
	}
	// create the directory that contains dependencies and script for the command
	subCommandDirName := getCommandDirAbsPath(command.Name)

	err := os.MkdirAll(subCommandDirName, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := cf.createEmptyShellScriptForCMD(command)
	if err != nil {
		return fmt.Errorf("Could not create empty Subcommand .sh file\n")
	}

	// write the contents of the command script to the persisted script
	_, err = blankFile.Write([]byte("hello"))
	if err != nil {
		return fmt.Errorf("Failed to copy the contennts of the target shell script\n")
	}

	return nil
}
