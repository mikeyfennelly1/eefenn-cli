package command_dir

import (
	"fmt"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func CreateSubcommandDirTree(commandId string) error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := GetSubcommandDependenciesDirectory(commandId)

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := CreateEmptySubcommandShellFile(commandId)
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
func GetAbsoluteSubcommandDirname(commandID string) string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIRoot, commandID)

	return commandDirectory
}

// GetSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func GetSubcommandDependenciesDirectory(commandId string) string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, commandId, commandId)

	return commandDependenciesDirectory
}

// CreateEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func CreateEmptySubcommandShellFile(commandId string) (*os.File, error) {
	fileName := GetSubcommandShellFileAbsPath(commandId)

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetSubcommandShellFileAbsPath(commandId string) string {
	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandId, commandId)

	return fileName
}

func removeCommandDirectoryRecursively(commandId string) error {
	dirname := GetAbsoluteSubcommandDirname(commandId)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}
