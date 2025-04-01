package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cli"
	"github.com/spf13/cobra"
	"os"
)

var (
	configFilePath string
	commandName    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ef", // The name of the command
	Short: "A command line tool for automating web application tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// command for committing an edited script to a command
var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit an edited script to a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.Commit(configFilePath)
		if err != nil {
			fmt.Printf("Unable to edit command '%s': %v\n", commandName, err)
		}
	},
}

func init() {
	// add flags to the 'ef commit' command
	commitCommand.Flags().StringVarP(&configFilePath, "config", "n", "", "The path the the config.yaml file for the command.")
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	rootCmd.AddCommand(commitCommand)
	err := commitCommand.MarkFlagRequired("config")
	if err != nil {
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
