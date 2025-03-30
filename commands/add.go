package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"github.com/eefenn/eefenn-cli/subcommand"
)

func Add(subcommand subcommand.Subcommand) error {
	fmt.Println("getting current config")
	config, err := config.GetCurrentConfig()
	if err != nil {
		return err
	}
	fmt.Println("Got current config")

	config.AddCommand(subcommand)
	config.Update()
	fmt.Println("Added command to config")

	err = command_dir.CreateSubcommandDirTree(subcommand.Hash)
	if err != nil {
		return err
	}
	fmt.Println("Created command directory tree")

	return nil
}
