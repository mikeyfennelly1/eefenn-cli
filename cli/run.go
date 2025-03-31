package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/utils"
	"os/exec"
)

func Run(commandName string, commandArgs []string) ([]byte, error) {
	// get the absolute path to the script
	scriptPath, err := utils.GetSubcommandShellFileAbsPath(commandName)
	if err != nil {
		return nil, err
	}

	// create a command object from the script
	cmd := exec.Command("sh", append([]string{*scriptPath}, commandArgs...)...)

	// run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Error running the command script: %v\n", err)
	}

	return output, nil
}
