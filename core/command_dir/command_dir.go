// command_dir.go
//
// For interacting with the directory tree that manages commands.
//
// Author: Mikey Fennelly <mikeyp.fennelly@gmail.com>

package command_dir

import (
	"fmt"
	cmd_config "github.com/eefenn/eefenn-cli/cmd"
	"io"
	"os"
	"path/filepath"
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
	CreateCMDDirTree(cmd cmd_config.Command) error

	// RemoveCommandDirectoryRecursively
	//
	// remove a command directory recursively by command hash.
	RemoveCommandDirectoryRecursively(commandName string) error
}

type EefennCLIDirectoryTree struct{}

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func (edt *EefennCLIDirectoryTree) CreateCMDDirTree(cmd cmd_config.Command) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must have root permissions to perform changes to CLI core\n")
	}
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := getCMDDependenciesDir(cmd.Name)

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := edt.createEmptyShellScriptForCMD(cmd)
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

// CreateEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func (edt *EefennCLIDirectoryTree) createEmptyShellScriptForCMD(cmd cmd_config.Command) (*os.File, error) {
	fileName := getSubcommandShellFileAbsPath(cmd.Name)

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
	dirname := getAbsoluteSubcommandDirname(commandName)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}

// CopyScriptToCMDDir
//
// Move a shell script to its command's directory
func (edt *EefennCLIDirectoryTree) CopyScriptToCMDDir(cmd cmd_config.Command) error {
	sourceFile, err := os.OpenFile(cmd.Script, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := edt.createEmptyShellScriptForCMD(cmd)
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

// CopyScriptToCMDDir
//
// Move a shell script to its command's directory
func (edt *EefennCLIDirectoryTree) CopyDependenciesToDependenciesDir(cmd cmd_config.Command) error {
	for _, dependency := range cmd.Needs {
		err := edt.copyDependencyFileToDependencyDir(dependency, cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (edt *EefennCLIDirectoryTree) copyDependencyFileToDependencyDir(dependencyPath string, cmd cmd_config.Command) error {
	dependencyFile, err := os.OpenFile(dependencyPath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer dependencyFile.Close()

	dependencyFileName := filepath.Base(dependencyPath)
	cmdDepsDir := getCMDDependenciesDir(cmd.Name)
	destinationFilePath := fmt.Sprintf("%s/%s", cmdDepsDir, dependencyFileName)

	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, dependencyFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync() // Ensure all writes are flushed to disk

}
