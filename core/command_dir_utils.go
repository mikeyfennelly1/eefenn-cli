package core

import (
	"fmt"
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

// getAbsImgDirPath
//
// get the absolute directory path for the Subcommand directory.
func getAbsImgDirPath(commandName string) string {
	// create the string for the command ID
	imgDir := fmt.Sprintf("%s/%s", EefennCLIRoot, commandName)

	return imgDir
}
