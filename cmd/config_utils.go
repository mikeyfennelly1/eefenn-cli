package cmd

import (
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

}

func getConfigYamlInPWD() (*os.File, error) {
	return nil, nil
}

// Returns an error if syntax of passed config file is invalid.
func validateConfigSyntax() error {
	return nil
}

// Returns an error if the directory tree is not valid for the command.
func validateDirectoryTreeWithConfig() error {

}
