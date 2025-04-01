package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core"
	"github.com/eefenn/eefenn-cli/core/core_utils"
	"os"
)

// Add a command by .yaml configuration file
func Commit() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	// construct the path for the configuration file
	configInPWD := fmt.Sprintf("%s/config.yaml", pwd)
	// Parse the command from a Yaml file at location 'filePath'
	thisCMD, err := cmd.ParseCommandFromYaml(configInPWD)
	if err != nil {
		return err
	}

	// check if the passed config data is valid
	cmdIsValid, err := configIsValid(*thisCMD)
	if !cmdIsValid || err != nil {
		return fmt.Errorf("Error parsing config: %s", err)
	}

	// If the thisCMD already exists, return an error
	if core_utils.CMDExists(thisCMD.Name) {
		return fmt.Errorf("Command already exists.\n")
	}

	core, err := core.GetCore()
	if err != nil {
		return err
	}

	// Add the command to the config file
	err = core.Config.AddCMD(*thisCMD)
	if err != nil {
		return err
	}
	// Create the directory tree for the command
	err = core.DirectoryTree.CreateCMDDirTree(*thisCMD)
	if err != nil {
		return err
	}
	// Copy the script for the command from the pwd to the script
	// in newly created directory tree.
	err = core.DirectoryTree.CopyScriptToCMDDir(*thisCMD)
	if err != nil {
		return err
	}

	err = core.DirectoryTree.CopyDependenciesToDependenciesDir(*thisCMD)
	if err != nil {
		return err
	}

	fmt.Printf("Added new command: %s\n", thisCMD.Name)
	return nil
}

func configIsValid(commandParsedFromConfig cmd.Command) (bool, error) {
	// loop through all dependency files for the command
	for _, file := range commandParsedFromConfig.Needs {
		// check if each file is found in the pwd or in the pwd subtree
		if fileFoundInPWDTree(file) == false {
			return false, fmt.Errorf("Could not find dependency file '%s'", file)
		}
	}
	return true, nil
}

func fileFoundInPWDTree(fileName string) bool {
	stat, _ := os.Stat(fileName)
	if stat != nil {
		return true
	} else {
		return false
	}
}
