package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core"
	"os"
)

// Add a command by .yaml configuration file
func Commit() error {
	// if there is a config file in the pwd, get its path
	pathToConfigInPWD, err := findConfigYamlInPWD()
	if err != nil {
		return err
	}
	fmt.Printf("Found config file '%s' in pwd.", pathToConfigInPWD)

	validateConfig(pathToConfigInPWD)

	// check if the passed config data is valid
	cmdIsValid := configFileDependenciesExistInPWD(*pCommandInPwd)
	if err != nil {
		return err
	}
	if !cmdIsValid {
		return fmt.Errorf("Invalid Command: %v\n", err)
	}

	// If the thisCMD already exists, return an error
	if core.cmdExists(thisCMD.Name) {
		return fmt.Errorf("Command already exists.\n")
	}

	currentCore, err := core.GetCore()
	if err != nil {
		return err
	}

	// commit the command to currentCore
	err = currentCore.Commit(*thisCMD)
	if err != nil {
		return err
	}

	fmt.Printf("Added new command: %s\n", thisCMD.Name)
	return nil
}

// findConfigYamlInPWD
// Get a yaml configuration file in pwd if one exists.
func findConfigYamlInPWD() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// construct the path for the configuration file
	return fmt.Sprintf("%s/config.yaml", pwd), nil
}

func configFileDependenciesExistInPWD(commandParsedFromConfig cmd.Command) bool {
	// loop through all dependency files for the command
	for _, file := range commandParsedFromConfig.Needs {
		// check if each file is found in the pwd or in the pwd subtree
		if fileFoundInPWDTree(file) == false {
			return false
		}
	}
	return true
}

func fileFoundInPWDTree(fileName string) bool {
	stat, _ := os.Stat(fileName)
	if stat != nil {
		return true
	} else {
		return false
	}
}

// validateConfig
// Takes the path to a config file and checks if the syntax is valid
func validateConfig(pathToConfigFile string) error {

}
