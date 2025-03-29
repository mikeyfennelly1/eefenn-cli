// subcommand.go
//
// asc (add subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package subcommand

import (
	"github.com/eefenn/eefenn-cli/cmd/subcommand/command_dir"
	"github.com/google/uuid"
)

type subcommand struct {
	// alias of the script
	Name string `json:"name"`

	// unique identifier for the command
	Hash uuid.UUID `json:"command-hash"`

	// the script which the command is an alias for
	SourceScript string `json:"script"`

	// description for what the script does
	Description string `json:"description,omitempty"`
}

// AddSubCommand
//
// Add a subcommand, and it's script to the user's CLI
func (sc *subcommand) AddSubCommand() error {
	// create directory structure
	err := command_dir.CreateSubcommandDirTree(sc.getSubcommandId())
	if err != nil {
		return err
	}

	// copy the shell script
	err = command_dir.CopyShellScript(sc.SourceScript, sc.getSubcommandId())
	if err != nil {
		return err
	}

	// update the eefenn-cli.config.json to contain the command info

	return nil
}

// CreateSubCommand
//
// Create a subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, description string) subcommand {
	UUID := uuid.New()
	subCommand := subcommand{
		Name:         name,
		Hash:         UUID,
		SourceScript: sourceScriptName,
		Description:  description,
	}
	return subCommand
}

func (sc *subcommand) getSubcommandId() string {
	return sc.Hash.String()[:8]
}
