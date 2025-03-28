package main

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ef", // The name of the command
	Short: "A command line tool for automating web application tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my CLI tool!")
	},
}

func main() {
	cmd.GreetCmd.Flags().StringP("name", "n", "World", "Name to greet")
	rootCmd.AddCommand(cmd.GreetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
