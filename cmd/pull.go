package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Pulls the latest commands and configurations from eefenn remote source of truth.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tidying...")
	},
}
