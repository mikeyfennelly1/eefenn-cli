package cli

import (
	"fmt"
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

	err = core.Commit(*pCMD)
	if err != nil {
		return err
	}

	fmt.Printf("Added command: %s\n", (*pCMD).Name)

	return nil
}
