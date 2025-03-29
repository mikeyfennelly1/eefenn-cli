// add_subcommand.go
//
// asc (add subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package add_subcommand

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const eefennCLIDir = "/usr/lib/eefenn-cli"
const configJSONPath = eefennCLIDir + "/eefenn-cli.config.json"

var AscCommand = &cobra.Command{
	Use:   "asc",
	Short: "Add a subcommand to eefenn-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding subcommand...")
	},
}

type subcommand struct {
	// alias of the script
	Name string `json:"name"`

	// unique identifier for the command
	Hash uuid.UUID `json:"command-hash"`

	// the script which the command is an alias for
	SourceScript string `json:"script"`

	// script dependencies
	Dependencies []string `json:"dependencies,omitempty"`

	// description for what the script does
	Description string `json:"description,omitempty"`
}

// AddSubCommand
//
// Add a subcommand, and it's script to the user's CLI
func AddSubCommand(name string, sourceScriptName string, dependencyPaths []string, description string) error {
	subCommand := createSubCommand(name, sourceScriptName, dependencyPaths, description)

	err := subCommand.updateConfig()
	if err != nil {
		return err
	}

	err = subCommand.createSubcommandDirTree()
	if err != nil {
		return err
	}

	return nil
}

// createSubCommand
//
// Create a subcommand struct based on required command information
func createSubCommand(name string, sourceScriptName string, dependencyPaths []string, description string) subcommand {
	UUID := uuid.New()
	subCommand := subcommand{
		Name:         name,
		Hash:         UUID,
		SourceScript: sourceScriptName,
		Dependencies: dependencyPaths,
		Description:  description,
	}
	return subCommand
}
