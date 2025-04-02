package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
	"os"
)

func Edit(commandName string) error {
	currentCore, err := core.GetCore()
	if err != nil {
		return err
	}

	err = currentCore.RecursivelyCopyCommandDirToPWD(commandName)
	if err != nil {
		return err
	}

	pwd, _ := os.Getwd()
	fmt.Printf("Command config copied to %s", pwd)

	return nil
}
