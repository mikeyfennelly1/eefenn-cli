package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
	"os"
)

func Edit(commandName string) error {
	pwd, _ := os.Getwd()
	fmt.Printf("Command config copied to %s", pwd)

	err := core.CreateCommandInDir(commandName, pwd)
	if err != nil {
		return err
	}

	return nil
}
