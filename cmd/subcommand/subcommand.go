// subcommand.go
//
// asc (add Subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package subcommand

import (
	"github.com/google/uuid"
)

const eefennCLIDir = "/usr/lib/eefenn-cli"
const configJSONPath = eefennCLIDir + "/eefenn-cli.config.json"

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
	err := sc.updateConfig()
	if err != nil {
		return err
	}

	err = sc.createSubcommandDirTree()
	if err != nil {
		return err
	}

	err = sc.copyShellFile()
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
