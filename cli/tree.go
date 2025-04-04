package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
	"os"
	"os/exec"
)

func Tree(commandName string) error {
	c, err := core.GetCore()
	if err != nil {
		return err
	}

	pCMD, _ := c.GetCommandByName(commandName)
	if pCMD == nil {
		return fmt.Errorf("the command %s does not exist", commandName)
	}

	commandDir := fmt.Sprintf("%s/%s/", core.EefennCLIRoot, pCMD.Name)
	command := fmt.Sprintf("tree %s", commandDir)

	cmd := exec.Command("sh", "-c", command)

	cmd.Stdout = os.Stdout

	cmd.Run()

	return nil
}
