package command_dir

import (
	"fmt"
	cmd_config "github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/eefenn/eefenn-cli/utils"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func CreateSubcommandDirTree(cmd cmd_config.Command) error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := GetSubcommandDependenciesDirectory(cmd)

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := CreateEmptySubcommandShellFile(cmd)
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
func GetAbsoluteSubcommandDirname(commandName string) string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIRoot, commandName)

	return commandDirectory
}

// GetSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func GetSubcommandDependenciesDirectory(cmd cmd_config.Command) string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, cmd.Name, cmd.Name)

	return commandDependenciesDirectory
}

// CreateEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func CreateEmptySubcommandShellFile(cmd cmd_config.Command) (*os.File, error) {
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
func RemoveCommandDirectoryRecursively(commandName string) error {
	dirname := GetAbsoluteSubcommandDirname(commandName)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}
