package cmd_config

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

	// the date that the command was updated
	DateCreated string `json:"dateCreated"`

	// parameters for the command
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// HasParameters
//
// check if a subcommand has parameters
func (sc *Subcommand) HasParameters() bool {
	if len(sc.Parameters) != 0 {
		return true
	} else {
		return false
	}
}

// List
//
// Print a subcommand in the format of the ef ls command
func (sc *Subcommand) List() {
	var hasParamsString string
	if sc.HasParameters() {
		hasParamsString = "true"
	} else {
		hasParamsString = "false"
	}
	fmt.Printf("%-10s %-30s %-30s\n", sc.Hash[:8], sc.Name, hasParamsString)
}

func (sc *Subcommand) getSubcommandId() string {
	return sc.Hash[:8]
}

// CreateSubCommand
//
// Create a Subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, description string) Subcommand {
	// parameter are initially empty
	params := []Parameter{}
	UUID := uuid.New().String()
	subCommand := Subcommand{
		Name:        name,
		Hash:        UUID,
		Script:      sourceScriptName,
		Description: description,
		Parameters:  params,
	}
	return subCommand
}
