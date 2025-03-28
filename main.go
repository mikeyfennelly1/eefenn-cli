package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd/asc"
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

var myCommand = &cobra.Command{
	Use:   "create",
	Short: "Creates an entry with a name, file, and description",
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
	myCommand.Flags().StringVarP(&name, "name", "n", "", "Name of the entity (required)")
	myCommand.Flags().StringVarP(&file, "file", "f", "", "Path to the file in the current directory (required)")
	myCommand.Flags().StringVarP(&description, "description", "d", "", "Description of the entity")

	// Mark flags as required
	err := myCommand.MarkFlagRequired("name")
	if err != nil {
		return
	}
	err = myCommand.MarkFlagRequired("file")
	if err != nil {
		return
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "app"}
	
	rootCmd.AddCommand(myCommand)

	rootCmd.AddCommand(asc.AscCommand)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
