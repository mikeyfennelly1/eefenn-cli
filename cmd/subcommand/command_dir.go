package subcommand

import (
	"fmt"
	"os"
)

// createSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the Subcommand
func (sc *Subcommand) createSubcommandDirTree() error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := sc.getSubcommandDependenciesDirectory()

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this Subcommand: %v\n", err)
	}

	subCommandDir := sc.getAbsoluteSubcommandDirname()
	// create a blank command script
	blankFile, err := sc.createEmptySubcommandShellFile(subCommandDir)
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

// getAbsoluteSubcommandDirname
//
// get the absolute directory path for the Subcommand directory.
func (sc *Subcommand) getAbsoluteSubcommandDirname() string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", eefennCLIDir, sc.Hash.String())

	return commandDirectory
}

// getSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func (sc *Subcommand) getSubcommandDependenciesDirectory() string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", eefennCLIDir, sc.Hash.String(), sc.Hash.String())

	return commandDependenciesDirectory
}

// createEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func (sc *Subcommand) createEmptySubcommandShellFile(parentDir string) (*os.File, error) {
	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s/%s%s", parentDir, sc.Hash.String(), ".sh")

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
