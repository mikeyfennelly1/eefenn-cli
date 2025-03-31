package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/eefenn/eefenn-cli/utils"
)

func Add(subcommand subcommand.Subcommand) error {
	// if the command already exists, return an error
	if utils.CommandExists(subcommand.Name) {
		return fmt.Errorf("Command already exists.\n")
	}

	// get a Config struct from the current config.json
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}
	currentConfig.AddCommand(subcommand)
	currentConfig.Update()

	// create the directory structure for the command
	err = command_dir.CreateSubcommandDirTree(subcommand.Hash)
	if err != nil {
		return err
	}

	// copy the shell script for the subcommand to the script
	// location in /usr/lib/eefenn-cli/<command_hash>
	err = command_dir.CopyShellScript(subcommand.Script, subcommand.Hash)
	if err != nil {
		return err
	}

	return nil
}
