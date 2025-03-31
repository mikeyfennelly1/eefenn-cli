package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"os"
)

func Edit(commandName string) error {
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	commandHash, err := currentConfig.GetCommandHash(commandName)
	if err != nil {
		return err
	}
	// get the path to the shell script for the current version of the command
	scriptPath := command_dir.GetSubcommandShellFileAbsPath(*commandHash)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	editFilePath := fmt.Sprintf(pwd + "/" + commandName + ".sh")
	os.Create(editFilePath)

	// get the content of the script of current version of the command
	currentScriptContent, err := os.ReadFile(scriptPath)
	if err != nil {
		return err
	}

	// write the content of the current script into the newly created file
	err = os.WriteFile(editFilePath, currentScriptContent, 0666)
	if err != nil {
		return err
	}

	return nil
}
