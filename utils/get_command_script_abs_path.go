package utils

import "fmt"

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// GetSubcommandShellFileAbsPath
//
// Get the absolute path to the shell script for the command
// based on commandHash
func GetSubcommandShellFileAbsPath(commandName string) string {
	// return '<command-name>.sh' filename string
	fileName := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandName, commandName)

	return fileName
}
