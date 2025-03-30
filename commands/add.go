package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"github.com/eefenn/eefenn-cli/subcommand"
)

func Add(subcommand subcommand.Subcommand) error {
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}

	// check if the command already exists
	commandHash, err := currentConfig.GetCommandHash(subcommand.Name)
	if commandHash != nil {
		return fmt.Errorf("Command already exists.\n")
	}

	currentConfig.AddCommand(subcommand)
	currentConfig.Update()

	err = command_dir.CreateSubcommandDirTree(subcommand.Hash)
	if err != nil {
		return err
	}

	return nil
}
