package subcommand

import "fmt"

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
	fmt.Printf("%-10s %-10s %-20s\n", sc.Hash[:8], sc.Name, sc.Description)
}

func (sc *Subcommand) getSubcommandId() string {
	return sc.Hash[:8]
}
