package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core"
	"github.com/eefenn/eefenn-cli/core/core_utils"
	"os"
)

// Add a command by .yaml configuration file
func Commit() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	configInPWD := fmt.Sprintf("%s/config.yaml", pwd)
	// Parse the command from a Yaml file at location 'filePath'
	cmd, err := cmd.ParseCommandFromYaml(configInPWD)
	if err != nil {
		return err
	}
	// If the cmd already exists, return an error
	if core_utils.CommandExists(cmd.Name) {
		return fmt.Errorf("Command already exists.\n")
	}

	core, err := core.GetCore()
	if err != nil {
		return err
	}

	// Add the command to the config file
	err = core.Config.AddCommand(*cmd)
	if err != nil {
		return err
	}
	// Create the directory tree for the command
	err = core.DirectoryTree.CreateCommandDirTree(*cmd)
	if err != nil {
		return err
	}
	// Copy the script for the command from the pwd to the script
	// in newly created directory tree.
	err = core.DirectoryTree.CopyScriptToCommandDirectory(*cmd)
	if err != nil {
		return err
	}

	fmt.Printf("Added new command: %s\n", cmd.Name)
	return nil
}
