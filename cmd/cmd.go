package cmd

type CommandInterface interface {
	// Run
	//
	// Execute the script for the command with arguments.
	Run(args []Arg)
}

type Command struct {
	// alias of the script
	Name string `json:"name"`

	// the script which the command is an alias for
	Script string `json:"script"`

	// array of filepaths that the script needs to run
	Needs []string `json:"needs"`

	// description for what the script does
	Description string `json:"description"`

	// the arguments to the command
	Args []Arg
}

type Arg struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
