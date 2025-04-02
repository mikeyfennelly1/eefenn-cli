package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configYamlPath = "./config.yaml"

// GetCMDFromPWD
// Looks in the pwd for a config file.
// Performs syntactic validation of the configuration file.
// Checks the directory to find dependencies that are specified
// in the file.
//
// Returns the appropriate error operation if this fails.
func GetCMDFromPWD() (*Command, error) {
	// ensure that the config.yaml exists in the pwd
	yamlContents, err := os.ReadFile(configYamlPath)
	if err != nil {
		return nil, err
	}

	// unmarshal a command from the config.yaml
	pCMD, err := unMarshalCommandFromYamlContents(yamlContents)
	if err != nil {
		return nil, err
	}

	// validate the syntax of the command
	err = validateCMDSyntax(pCMD)
	if err != nil {
		return nil, err
	}

	// ensure that script and its dependencies exist at
	// specified locations.
	err = validateDirectoryTreeWithCMD(pCMD)

	return pCMD, nil
}

// Returns an error if the directory tree is not valid for the command.
func validateDirectoryTreeWithCMD(cmd *Command) error {
	// ensure script is findable at specified
	// position relative to script
	_, err := os.Stat(cmd.Script)
	if err != nil {
		return err
	}

	for _, dependency := range cmd.Needs {
		// check that all dependencies exist
		_, err := os.Stat(dependency)
		if err != nil {
			return err
		}
	}

	return nil
}

func unMarshalCommandFromYamlContents(yamlConfigContents []byte) (*Command, error) {
	var cmd Command
	err := yaml.Unmarshal(yamlConfigContents, &cmd)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return nil, err
	}

	return &cmd, nil
}

// Returns an error if syntax of passed
// command parsed from the config is invalid.
func validateCMDSyntax(cmd *Command) error {
	// check for empty string as command name
	if !(len(cmd.Name) > 0) {
		return fmt.Errorf("invalid command name")
	}

	// check for empty command description
	if !(len(cmd.Description) > 0) {
		return fmt.Errorf("invalid command description")
	}

	if cmd.Args != nil {
		for _, arg := range cmd.Args {
			// check for empty argument name string
			if !(len(arg.Name) > 0) {
				return fmt.Errorf("invalid argument name")
			}
			// check for empty argument description string
			if !(len(arg.Description) > 0) {
				return fmt.Errorf("invalid argument description")
			}
			// check for invalid type for argument
			switch arg.Type {
			case "string":
				continue
			case "int":
				continue
			default:
				return fmt.Errorf("unrecognised type '%s' for argument '%s'", arg.Type, arg.Name)
			}
		}
	}
	return nil
}
