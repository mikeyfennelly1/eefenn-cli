package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cli"
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/spf13/cobra"
	"os"
)

var (
	file          string
	name          string
	description   string
	commandName   string
	commitMessage string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ef", // The name of the command
	Short: "A command line tool for automating web application tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// command for adding a subcommand to eefenn-cli
var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a subcommand to eefenn-cli",
	Run: func(cmd *cobra.Command, args []string) {
		// Validate that the file exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist:", file)
			return
		}

		thisSubcommand := cmd_config.CreateSubCommand(name, file, description)
		err := cli.Add(thisSubcommand)
		if err != nil {
			fmt.Printf("Could not create subcommand: %v", err)
			return
		}
		fmt.Println(thisSubcommand.Hash)
	},
}

// command for listing subcommands in the binary
var lsCommand = &cobra.Command{
	Use:   "ls",
	Short: "List all subcommands",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.LS()
		if err != nil {
			return
		}
	},
}

// command for removing a command
var rmCommand = &cobra.Command{
	Use:   "rm",
	Short: "Remove a command",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.RemoveSubcommand(commandName)
		if err != nil {
			return
		}
	},
}

// command for running a command
var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run an eefenn-cli command",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(commandName, args)
		if err != nil {
			fmt.Printf("Unable to run command '%s': %v\n", commandName, err)
		}

		fmt.Printf(string(result))
	},
}

// command for editing an existing command's script
var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edit the script for a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.Edit(commandName)
		if err != nil {
			fmt.Printf("Unable to edit command '%s': %v\n", commandName, err)
			return
		}

		fmt.Printf("Created file to edit: %s.sh\n", commandName)
		fmt.Println("Edit the script, and use the 'commit' command to update the command")
	},
}

// command for committing an edited script to a command
var commitCommand = &cobra.Command{
	Use:   "commit",
	Short: "Commit an edited script to a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.Commit(commandName, commitMessage)
		if err != nil {
			fmt.Printf("Unable to edit command '%s': %v\n", commandName, err)
		}
	},
}

// command for printing the description of an existing command
var describeCommand = &cobra.Command{
	Use:   "describe",
	Short: "Print the description of a command.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.Describe(commandName)
		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}

// command for printing the description of an existing command
var addDescriptionCommand = &cobra.Command{
	Use:   "add-description",
	Short: "Add description to the command",
	Run: func(cmd *cobra.Command, args []string) {
		err := cli.AddDescription(commandName, description)
		if err != nil {
			fmt.Printf("Could not add description: %v\n", err)
		}
	},
}

func init() {
	// add flags to the 'ef add' command
	addCommand.Flags().StringVarP(&name, "name", "n", "", "Name of the entity (required)")
	addCommand.Flags().StringVarP(&file, "file", "f", "", "Path to the file in the current directory (required)")
	addCommand.Flags().StringVarP(&description, "description", "d", "", "Description of what the command does.")
	// add flags to the 'ef rm' command
	rmCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to remove.")
	// add flags to the 'ef run' command
	runCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to run.")

	// add flags to the 'ef edit' command
	editCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to edit.")

	// add flags to the 'ef describe' command
	describeCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to run.")

	// add flags to the 'ef commit' command
	commitCommand.Flags().StringVarP(&commandName, "name", "n", "", "The name of the command you want to commit.")
	commitCommand.Flags().StringVarP(&commitMessage, "message", "m", "", "The commit message for the commit.")

	// add flags to the 'ef commit' command
	addDescriptionCommand.Flags().StringVarP(&description, "new-description", "nd", "", "The new description that you want to add.")
}

func main() {
	// ensure that binary is running with root permissions before running
	if os.Geteuid() != 0 {
		fmt.Println("You must be superuser to run this binary.")
		return
	}

	// add commands to root command
	rootCmd.AddCommand(addCommand)
	// specify required flags for the 'ef add' command
	err := addCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = addCommand.MarkFlagRequired("file")
	if err != nil {
		return
	}
	err = addCommand.MarkFlagRequired("description")
	if err != nil {
		return
	}

	rootCmd.AddCommand(lsCommand)
	rootCmd.AddCommand(rmCommand)
	err = rmCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}

	rootCmd.AddCommand(runCommand)
	rootCmd.AddCommand(describeCommand)
	err = describeCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}

	rootCmd.AddCommand(editCommand)
	err = editCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}

	rootCmd.AddCommand(commitCommand)
	err = commitCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = commitCommand.MarkFlagRequired("message")
	if err != nil {
		return
	}

	rootCmd.AddCommand(addDescriptionCommand)
	err = commitCommand.MarkFlagRequired("new-description")
	if err != nil {
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
