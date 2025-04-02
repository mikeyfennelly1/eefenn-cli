package cli

import (
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core"
)

// Commit
//
// Add a command by .yaml configuration file
func Commit() error {
	pCMD, err := cmd.GetCMDFromPWD()
	if err != nil {
		return err
	}

	currentCore, err := core.GetCore()
	if err != nil {
		return err
	}

	err = currentCore.Commit(*pCMD)
	if err != nil {
		return err
	}

	return nil
}
