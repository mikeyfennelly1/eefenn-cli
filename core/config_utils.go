package core

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
func GetCMDFromPWD() (*CommandInterface, error) {
	// ensure that the config.yaml exists in the pwd
	yamlContents, err := os.ReadFile(configYamlPath)
	if err != nil {
		return nil, err
	}

	// unmarshal a CommandImage from the config.yaml
	pCMD, err := unMarshalCommandFromYamlContents(yamlContents)
	if err != nil {
		return nil, err
	}

	// validate the syntax of the CommandImage
	err = validateCMDSyntax(pCMD)
	if err != nil {
		return nil, err
	}

	// ensure that script and its dependencies exist at
	// specified locations.
	err = validateDirectoryTreeWithCMD(pCMD)

	return pCMD, nil
}

// Returns an error if the directory tree is not valid for the CommandImage.
func validateDirectoryTreeWithCMD(cmd *CommandInterface) error {
	// ensure script is findable at specified
	// position relative to script
	_, err := os.Stat(cmd.Script)
	if err != nil {
		return err
	}

	for _, dependency := range cmd.Dependencies {
		// check that all dependencies exist
		_, err := os.Stat(dependency)
		if err != nil {
			return err
		}
	}

	return nil
}

func unMarshalCommandFromYamlContents(yamlConfigContents []byte) (*CommandInterface, error) {
	var cmd CommandInterface
	err := yaml.Unmarshal(yamlConfigContents, &cmd)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return nil, err
	}

	return &cmd, nil
}

// Returns an error if syntax of passed
// CommandImage parsed from the config is invalid.
func validateCMDSyntax(cmd *CommandInterface) error {
	// check for empty string as CommandImage name
	if !(len(cmd.Name) > 0) {
		return fmt.Errorf("invalid CommandImage name")
	}

	// check for empty CommandImage description
	if !(len(cmd.Description) > 0) {
		return fmt.Errorf("invalid CommandImage description")
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
