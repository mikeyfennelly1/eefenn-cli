package cmd_config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Command struct {
	// array of filepaths that the script needs to run
	Needs []string `yaml:"needs"`

	// alias of the script
	Name string `yaml:"name"`

	// the script which the command is an alias for
	Script string `yaml:"script"`

	// description for what the script does
	Description string `yaml:"description"`

	// the arguments to the command
	Args []struct {
		Name        string `yaml:"string"`
		Description string `yaml:"string"`
	}
}

func GetCommandFromYml() {
	yamlData, err := os.ReadFile("config.yaml")

	var config Command
	err = yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return
	}

	// Print parsed struct
	fmt.Printf("Parsed struct: %+v\n", config)

}
