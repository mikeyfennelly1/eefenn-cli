package cmd

type CommandInterface interface {
	GetCommandFilePathsRelativeToCMD() ([]string, error)
}

type Command struct {
	// alias of the script
	Name string `yaml:"name" json:"name"`

	// the script which the command is an alias for
	Script string `yaml:"script" json:"script"`

	// array of filepaths that the script needs to run
	Dependencies []string `yaml:"dependencies" json:"dependencies"`

	// description for what the script does
	Description string `yaml:"description" json:"description"`

	// the arguments to the command
	Args []Arg ` yaml:"args" json:"args"`
}

type Arg struct {
	Name        string `json:"name" json:"name"`
	Type        string `json:"type" json:"type"`
	Description string `json:"description" json:"description"`
}

// GetCommandFilePaths
// Gets the file paths for script, configuration file and all dependencies
// and returns these paths relative to where the command currently is.
func (cmd *Command) GetCommandFilePathsRelativeToCMD() ([]string, error) {
	var commandFilePaths []string

	// add command script
	commandFilePaths = append(commandFilePaths, cmd.Script)

	// add config.yaml

	// add all dependency files
	for _, dependency := range cmd.Dependencies {
		commandFilePaths = append(commandFilePaths, dependency)
	}

	return nil, nil
}
