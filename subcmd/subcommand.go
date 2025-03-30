// Subcommand.go
//
// asc (add Subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package subcmd

import (
	command_dir2 "github.com/eefenn/eefenn-cli/command_dir"
	"github.com/google/uuid"
)

type Subcommand struct {
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
// Add a Subcommand, and it's script to the user's CLI
func (sc *Subcommand) AddSubCommand() error {
	// create directory structure
	err := command_dir2.CreateSubcommandDirTree(sc.Hash.String())
	if err != nil {
		return err
	}

	// copy the shell script
	err = command_dir2.CopyShellScript(sc.SourceScript, sc.Hash.String())
	if err != nil {
		return err
	}

	// update the eefenn-cli.config.json to contain the command info

	return nil
}

// RemoveSubcommand
//
// Remove a subcommand's directories by command hash
func RemoveSubcommand(commandHash string, commandName string) error {
	err := command_dir2.RemoveCommandDirectoryRecursively(commandHash)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// CreateSubCommand
//
// Create a Subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, description string) Subcommand {
	UUID := uuid.New()
	subCommand := Subcommand{
		Name:         name,
		Hash:         UUID,
		SourceScript: sourceScriptName,
		Description:  description,
	}
	return subCommand
}

func (sc *Subcommand) getSubcommandId() string {
	return sc.Hash.String()[:8]
}
