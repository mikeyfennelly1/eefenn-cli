package subcommand

import (
	"fmt"
	"github.com/google/uuid"
)

type Subcommand struct {
	// alias of the script
	Name string `json:"name"`

	// unique identifier for the command
	Hash string `json:"hash"`

	// the script which the command is an alias for
	Script string `json:"script"`

	// description for what the script does
	Description string `json:"desc,omitempty"`

	DateCreated string `json:"dateCreated"`
}

// List
//
// Print a subcommand in the format of the ef ls command
func (sc *Subcommand) List() {
	fmt.Printf("%-10s %-10s\n", sc.Hash[:8], sc.Name)
}

func (sc *Subcommand) getSubcommandId() string {
	return sc.Hash[:8]
}

// CreateSubCommand
//
// Create a Subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, description string) Subcommand {
	UUID := uuid.New().String()
	subCommand := Subcommand{
		Name:        name,
		Hash:        UUID,
		Script:      sourceScriptName,
		Description: description,
	}
	return subCommand
}
