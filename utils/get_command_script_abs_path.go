package utils

import "fmt"

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// GetSubcommandShellFileAbsPath
//
// Get the absolute path to the shell script for the command
// based on commandHash
func GetSubcommandShellFileAbsPath(commandName string) (*string, error) {
	commandHash, err := GetCommandHash(commandName)
	if err != nil {
		return nil, err
	}

	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandHash, commandHash)

	return &fileName, err
}
