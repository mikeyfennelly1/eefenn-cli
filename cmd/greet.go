package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GreetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greets a person",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Printf("Hello, %s!\n", name)
	},
}
