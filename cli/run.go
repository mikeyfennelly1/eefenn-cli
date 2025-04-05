package cli

import (
	"github.com/eefenn/eefenn-cli/core"
	"os"
)

func Run(cmdName string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = core.Run(cmdName, pwd)
	if err != nil {
		return err
	}
	
	return nil
}
