package main

import (
	"fmt"
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

		// Print the collected data
		fmt.Println("Name:", name)
		fmt.Println("File:", file)
		fmt.Println("Description:", description)
	},
}

func init() {
	ascCommand.Flags().StringVarP(&name, "name", "n", "", "Name of the entity (required)")
	ascCommand.Flags().StringVarP(&file, "file", "f", "", "Path to the file in the current directory (required)")
	ascCommand.Flags().StringVarP(&description, "description", "d", "", "Description of the entity")

	// Mark flags as required
	err := ascCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = ascCommand.MarkFlagRequired("file")
	if err != nil {
		return
	}
}

func main() {
	rootCmd.AddCommand(ascCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
