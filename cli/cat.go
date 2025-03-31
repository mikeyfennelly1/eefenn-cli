package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/utils"
)

// Cat
//
// print the contents of the script file that the command
// is an alias for
func Cat(commandName string) error {
	// get the contents of the commands script
	scriptContents, err := utils.GetScriptContents(commandName)
	if err != nil {
		return err
	}

	// print the contents
	fmt.Printf("%s\n", string(scriptContents))

	return nil
}
