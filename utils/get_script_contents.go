package utils

import (
	"os"
)

func GetScriptContents(commandName string) ([]byte, error) {
	// get the Subcommand structre for <commandName>
	_, command, err := GetCommand(commandName)
	if err != nil {
		return nil, err
	}

	// get the absolute path of the script for the command
	commandScriptAbsPath, err := GetSubcommandShellFileAbsPath(command.Hash)
	if err != nil {
		return nil, err
	}

	// read the contents of the script into commandScriptContents
	commandscriptContents, err := os.ReadFile(*commandScriptAbsPath)
	if err != nil {
		return nil, err
	}

	return commandscriptContents, nil
}
