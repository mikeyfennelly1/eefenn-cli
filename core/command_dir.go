// command_dir.go
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
	// CreateCMDDirTree
	//
	// Create an entry in /usr/lib/eefenn-cli for the Subcommand.
	createCMDDir(command cmd.Command) error

	// RemoveCommandDirectoryRecursively
	//
	// remove a command directory recursively by command hash.
	RemoveCommandDirectoryRecursively(commandName string) error
}

type eefennCLIDirectoryTree struct{}

// CreateCMDDir
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func (edt *eefennCLIDirectoryTree) createCMDDir(command cmd.Command) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must have root permissions to perform changes to CLI core\n")
	}
	// create the directory that contains dependencies and script for the command
	cmdImageDir := getAbsImgDirPath(command.Name)
	err := os.MkdirAll(cmdImageDir, 0755)
	if err != nil {
		return fmt.Errorf("could not create directory for this command: %v\n", err)
	}

	return nil
}

func (edt *eefennCLIDirectoryTree) CopyCommandFilesToCMDDir(command cmd.Command) error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must have root permissions to perform changes to CLI core\n")
	}
	// create the directory that contains dependencies and script for the command
	filesNeeded := command.GetCmdFilePaths()

	cmdDir := command.GetCmdDir()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = copyFile(pwd, cmdDir)
	if err != nil {
		return err
	}

	return nil
}

// RemoveCommandDirectoryRecursively
//
// remove a command directory recursively by command hash
func (edt *eefennCLIDirectoryTree) RemoveCommandDirectoryRecursively(commandName string) error {
	dirname := getAbsImgDirPath(commandName)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}
