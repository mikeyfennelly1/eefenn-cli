package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/commands"
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/spf13/cobra"
	"os"
)

var (
	file        string
	name        string
	description string
	commandName string
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
	Use:   "add",
	Short: "Add a subcommand to eefenn-cli",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate that the file exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist:", file)
			return
		}

		thisSubcommand := subcommand.CreateSubCommand(name, file, description)
		err := commands.Add(thisSubcommand)
		if err != nil {
			fmt.Printf("Could not create subcommand: %v", err)
			return
		}
		fmt.Println(thisSubcommand.Hash)
	},
}

var lsCommand = &cobra.Command{
	Use:   "ls",
	Short: "List all subcommands",
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.LS()
		if err != nil {
			return
		}
	},
}

var rmCommand = &cobra.Command{
	Use:   "rm",
	Short: "Remove a command",
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.RemoveSubcommand(commandName)
		if err != nil {
			return
		}
	},
}

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run an eefenn-cli command",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := commands.Run(commandName)
		if err != nil {
			fmt.Printf("Unable to run command '%s': %v\n", commandName, err)
		}

		fmt.Printf(string(result))
	},
}

var describeCommand = &cobra.Command{
	Use:   "describe",
	Short: "Print the description of a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.Describe(commandName)
		if err != nil {
			fmt.Printf("%v", err)
		}

	},
}

func init() {
	ascCommand.Flags().StringVarP(&name, "name", "n", "", "Name of the entity (required)")
	ascCommand.Flags().StringVarP(&file, "file", "f", "", "Path to the file in the current directory (required)")
	ascCommand.Flags().StringVarP(&description, "description", "d", "", "Description of what the command does.")

	rmCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to remove.")

	runCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to run.")

	describeCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to run.")

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

	err = rmCommand.MarkFlagRequired("name")
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

	rootCmd.AddCommand(lsCommand)

	rootCmd.AddCommand(rmCommand)

	rootCmd.AddCommand(runCommand)

	rootCmd.AddCommand(describeCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
