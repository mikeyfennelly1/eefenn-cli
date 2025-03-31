package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/utils"
	"os"
)

// Edit
//
// Copies the current script of the command into users pwd.
// user can then make changes to this script and commit these
// to the command via 'ef commit'
func Edit(commandName string) error {
	scriptContents, err := utils.GetScriptContents(commandName)
	if err != nil {
		return err
	}

	// get the pwd of where the edit command was ran
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	editFilePath := fmt.Sprintf(pwd + "/" + commandName + ".sh")
	// create the file to copy the current script contents to
	os.Create(editFilePath)

	// write the content of the current script into the newly created file
	err = os.WriteFile(editFilePath, scriptContents, 0666)
	if err != nil {
		return err
	}

	return nil
}
