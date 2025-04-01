// command_dir.go
//
// For interacting with the directory tree that manages commands.
//
// Author: Mikey Fennelly <mikeyp.fennelly@gmail.com>

package command_dir

import (
	"fmt"
	cmd_config "github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/eefenn/eefenn-cli/core/config"
	"github.com/eefenn/eefenn-cli/utils"
	"io"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// EefennCLIDirectoryTreeInterface
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
type EefennCLIDirectoryTreeInterface interface {
	// CreateSubcommandDirTree
	//
	// Create an entry in /usr/lib/eefenn-cli for the Subcommand.
	CreateSubcommandDirTree(cmd cmd_config.Command) error

	// GetAbsoluteSubcommandDirname
	//
	// get the absolute directory path for the Subcommand directory.
	GetAbsoluteSubcommandDirname(commandName string) string

	// GetSubcommandDependenciesDirectory
	//
	// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
	GetSubcommandDependenciesDirectory(cmd cmd_config.Command) string

	// CreateEmptySubcommandShellFile
	//
	// Create an empty shell file of the name <command-hash>.sh
	CreateEmptySubcommandShellFile(cmd cmd_config.Command) (*os.File, error)

	// RemoveCommandDirectoryRecursively
	//
	// remove a command directory recursively by command hash.
	RemoveCommandDirectoryRecursively(commandName string) error
}

type EefennCLIDirectoryTree struct {
	Config config.Config
}

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func (edt *EefennCLIDirectoryTree) CreateSubcommandDirTree(cmd cmd_config.Command) error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := edt.GetSubcommandDependenciesDirectory(cmd)

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := edt.CreateEmptySubcommandShellFile(cmd)
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

// GetAbsoluteSubcommandDirname
//
// get the absolute directory path for the Subcommand directory.
func (edt *EefennCLIDirectoryTree) GetAbsoluteSubcommandDirname(commandName string) string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIRoot, commandName)

	return commandDirectory
}

// GetSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func (edt *EefennCLIDirectoryTree) GetSubcommandDependenciesDirectory(cmd cmd_config.Command) string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, cmd.Name, cmd.Name)

	return commandDependenciesDirectory
}

// CreateEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func (edt *EefennCLIDirectoryTree) CreateEmptySubcommandShellFile(cmd cmd_config.Command) (*os.File, error) {
	fileName := utils.GetSubcommandShellFileAbsPath(cmd.Name)

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// RemoveCommandDirectoryRecursively
//
// remove a command directory recursively by command hash
func (edt *EefennCLIDirectoryTree) RemoveCommandDirectoryRecursively(commandName string) error {
	dirname := edt.GetAbsoluteSubcommandDirname(commandName)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}

func (edt *EefennCLIDirectoryTree) CopyShellScript(cmd cmd_config.Command) error {
	sourceFile, err := os.OpenFile(cmd.Script, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := edt.CreateEmptySubcommandShellFile(cmd)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync() // Ensure all writes are flushed to disk
}
