package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/eefenn/eefenn-cli/core"
	"github.com/eefenn/eefenn-cli/core/config"
	"github.com/eefenn/eefenn-cli/utils"
)

func Add(cmd cmd_config.Command) error {
	// if the cmd already exists, return an error
	if utils.CommandExists(cmd.Name) {
		return fmt.Errorf("Command already exists.\n")
	}

	// get a Config struct from the current config.json
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}
	err = currentConfig.AddCommand(cmd)
	if err != nil {
		return err
	}
	err = currentConfig.Update()
	if err != nil {
		return err
	}

	// create the directory structure for the cmd
	err = command_dir.CreateSubcommandDirTree(cmd)
	if err != nil {
		return err
	}

	// copy the shell script for the subcommand to the script
	// location in /usr/lib/eefenn-cli/<command_hash>
	err = command_dir.CopyShellScript(cmd)
	if err != nil {
		return err
	}

	return nil
}
