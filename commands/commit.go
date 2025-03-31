package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"os"
)

func Commit(commandName string, commitMessage string) error {
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	editFile := fmt.Sprintf("%s/%s%s", pwd, commandName, ".sh")
	fileInfo, _ := os.Stat(editFile)
	if fileInfo == nil {
		return fmt.Errorf("Could not find file '%s.sh' in current directory\n.", commandName)
	}

	commandHash, err := currentConfig.GetCommandHash(commandName)
	if err != nil {
		return err
	}
	commandScriptAbsPath := command_dir.GetSubcommandShellFileAbsPath(*commandHash)

	updatedFile, err := os.ReadFile(editFile)
	if err != nil {
		return err
	}
	err = os.WriteFile(commandScriptAbsPath, updatedFile, 0666)
	if err != nil {
		return err
	}

	fmt.Println(commitMessage)

	return nil
}
