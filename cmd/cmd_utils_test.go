package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// GetCommandFromYml
//
// Create a cmd object from a .yaml configuration file
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
