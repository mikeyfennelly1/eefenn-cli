package cmd

type CommandInterface interface {
	// Run
	//
	// Execute the script for the command with arguments.
	Run(args []Arg)

	GetName() string

	GetScriptName() string

	GetDependencies() []string

	GetDescription() string

	ToJSON() []byte

	ToYAML() []byte
}

type CommandYaml struct {
	// alias of the script
	Name string `yaml:"name"`

	// the script which the command is an alias for
	Script string `yaml:"script"`

	// array of filepaths that the script needs to run
	Dependencies []string `yaml:"dependencies"`

	// description for what the script does
	Description string `yaml:"description"`

	// the arguments to the command
	Args []Arg
}

func (c CommandYaml) ToJSON() []byte {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) ToYAML() []byte {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) Run(args []Arg) {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) GetScriptName() string {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) GetDependencies() []string {
	//TODO implement me
	panic("implement me")
}

func (c CommandYaml) GetDescription() string {
	//TODO implement me
	panic("implement me")
}

type Arg struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
