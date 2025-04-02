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
	Short: "A command line tool for automating shell tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// command for committing an edited script to a command
var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit an edited script to a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.Commit()
		if err != nil {
			fmt.Printf("Error committing command: %v\n", err)
		}
	},
}

// command for committing an edited script to a command
var rmCommand = &cobra.Command{
	Use:   "rm",
	Short: "Remove a command by name.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.RM(commandName)
		if err != nil {
			fmt.Printf("Error removing command: %v\n", err)
		}
	},
}

func init() {
	rmCommand.Flags().StringVarP(&commandName, "name", "n", "", "User name (required)")
	err := rmCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	rootCmd.AddCommand(commitCommand)

	rootCmd.AddCommand(rmCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
