package core

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
)

// getSubcommandShellFileAbsPath
//
// Get the absolute path to the shell script for the command
// based on commandHash
func getSubcommandShellFileAbsPath(commandName string) string {
	// return '<command-name>.sh' filename string
	fileName := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandName, commandName)

	return fileName
}

// getAbsoluteSubcommandDirname
//
// get the absolute directory path for the Subcommand directory.
func getAbsoluteSubcommandDirname(commandName string) string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIRoot, commandName)

	return commandDirectory
}

// getCMDDependenciesDir
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func getCMDDependenciesDir(commandName string) string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, commandName, commandName)

	return commandDependenciesDirectory
}

func changeCommandDirectory(command cmd.Command) {

}
