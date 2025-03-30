package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/command_dir"
	"github.com/eefenn/eefenn-cli/config"
	"os/exec"
)

func Run(commandName string) ([]byte, error) {
	// get the current config
	currentConfig, err := config.GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	// get the hash for the command in eefenn-cli.config.json
	hash, err := currentConfig.GetCommandHash(commandName)
	if err != nil {
		return nil, fmt.Errorf("Could not find hash for command: %v\n", err)

	}

	// get the absolute path to the script
	scriptPath := command_dir.GetSubcommandShellFileAbsPath(*hash)

	// create a command object from the script
	cmd := exec.Command("sh", scriptPath)

	// run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Error running the command script: %v\n", err)
	}

	return output, nil
}
