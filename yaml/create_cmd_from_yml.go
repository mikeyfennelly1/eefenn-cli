package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Command struct {
	// alias of the script
	Name string `yaml:"name"`

	// the script which the command is an alias for
	Script string `yaml:"script"`

	// array of filepaths that the script needs to run
	Needs []string `yaml:"needs"`

	// description for what the script does
	Description string `yaml:"description"`

	// the arguments to the command
	Args []Arg
}

type Arg struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

func GetCommandFromYml(filePath string) (*Command, error) {
	yamlData, err := os.ReadFile(filePath)

	var cmd Command
	err = yaml.Unmarshal([]byte(yamlData), &cmd)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return nil, err
	}

	return &cmd, nil
}
