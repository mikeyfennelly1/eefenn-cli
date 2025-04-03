package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
)

// Commit
//
// Add a command by .yaml configuration file
func Commit() error {
	pCMD, err := core.GetCMDFromPWD()
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

	fmt.Printf("Added command: %s\n", (*pCMD).Name)

	return nil
}
