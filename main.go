package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd/add_subcommand"
	"github.com/spf13/cobra"
	"os"
)

var (
	file        string
	name        string
	description string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ef", // The name of the command
	Short: "A command line tool for automating web application tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: ef <subcommand>")
	},
}

var ascCommand = &cobra.Command{
	Use:   "asc",
	Short: "Add a subcommand to eefenn-cli",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate that the file exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist:", file)
			return
		}

		subCommand := add_subcommand.CreateSubCommand(name, file, description)
		err := subCommand.AddSubCommand()
		if err != nil {
			fmt.Printf("Could not create subcommand %v", err)
			return
		}
		fmt.Println(subCommand.Hash.String())
	},
}

func init() {
	ascCommand.Flags().StringVarP(&name, "name", "n", "", "Name of the entity (required)")
	ascCommand.Flags().StringVarP(&file, "file", "f", "", "Path to the file in the current directory (required)")
	ascCommand.Flags().StringVarP(&description, "description", "d", "", "Description of what the command does.")

	// Mark flags as required
	err := ascCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = ascCommand.MarkFlagRequired("file")
	if err != nil {
		return
	}
	err = ascCommand.MarkFlagRequired("description")
	if err != nil {
		return
	}
}

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}
	rootCmd.AddCommand(ascCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
