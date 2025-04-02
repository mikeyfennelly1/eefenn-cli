package cmd

import (
	"fmt"
	"os"
)

// GetCMDFromPWD
// Looks in the pwd for a config file.
// Performs syntactic validation of the configuration file.
// Checks the directory to find dependencies that are specified
// in the file.
//
// Returns the appropriate error operation if this fails.
func GetCMDFromPWD() (*Command, error) {
	// get the config.yaml in the pwd.
	pYAML, err := getConfigYamlInPWD()
	if err != nil {
		return nil, err
	}

	// unmarshal a command from the config.yaml
	pCMD, err := unMarshalCommandFromYaml(pYAML)
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

func getConfigYamlInPWD() (*os.File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configYamlPath := fmt.Sprintf("%s/%s", pwd, "config.yaml")
	configYamlFile, err := os.OpenFile(configYamlPath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer func(configYamlFile *os.File) {
		err := configYamlFile.Close()
		if err != nil {
			panic("Error closing config.yaml. Exiting...")
		}
	}(configYamlFile)

	return configYamlFile, nil
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
		// check that all dependencies are findable
		// at respective positions relative to script
		_, err := os.Stat(dependency)
		if err != nil {
			return err
		}
	}

	return nil
}

func unMarshalCommandFromYaml(yamlConfig *os.File) (*Command, error) {
	return nil, nil
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
			switch arg.Name {
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
