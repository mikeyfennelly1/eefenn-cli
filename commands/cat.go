package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"os"
)

// Cat
//
// print the contents of the script file that the command
// is an alias for
func Cat(commandName string) error {
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	hash, err := currentConfig.GetCommandHash(commandName)
	if err != nil {
		return err
	}

	scriptAbsPath := command_dir.GetSubcommandShellFileAbsPath(*hash)
	scriptContents, err := os.ReadFile(scriptAbsPath)
	if err != nil {
		return err
	}

	fmt.Printf(string(scriptContents))

	return nil
}
