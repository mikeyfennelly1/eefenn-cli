package command_dir

import (
	"fmt"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func CreateSubcommandDirTree(commandHash string) error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := GetSubcommandDependenciesDirectory(commandHash)

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := CreateEmptySubcommandShellFile(commandHash)
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
func GetAbsoluteSubcommandDirname(commandHash string) string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIRoot, commandHash)

	return commandDirectory
}

// GetSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func GetSubcommandDependenciesDirectory(commandHash string) string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, commandHash, commandHash)

	return commandDependenciesDirectory
}

// CreateEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func CreateEmptySubcommandShellFile(commandHash string) (*os.File, error) {
	fileName := GetSubcommandShellFileAbsPath(commandHash)

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetSubcommandShellFileAbsPath(commandHash string) string {
	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandHash, commandHash)

	return fileName
}

func RemoveCommandDirectoryRecursively(commandHash string) error {
	dirname := GetAbsoluteSubcommandDirname(commandHash)

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}

	return nil
}
