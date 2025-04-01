package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/utils"
	"os"
)

func Commit(commandName string, commitMessage string) error {
	// get the contents of the updated script
	updatedFileContents, err := getUpdatedFileContents(commandName)
	if err != nil {
		return err
	}

	// get the absolute path of the script file for the command
	commandScriptAbsPath := utils.GetSubcommandShellFileAbsPath(commandName)
	if err != nil {
		return err
	}
	// write the contents of the updated file to the script for the command
	err = os.WriteFile(commandScriptAbsPath, updatedFileContents, 0666)
	if err != nil {
		return err
	}

	// only printing commit message for now
	fmt.Println(commitMessage)

	return nil
}

// Get contents of the updated file.
// File must have the same name as command
func getUpdatedFileContents(commandName string) ([]byte, error) {
	// get the pwd of where the 'commit' command was ran
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// path to the updated file
	updatedFilePath := fmt.Sprintf("%s/%s%s", pwd, commandName, ".sh")
	fileInfo, _ := os.Stat(updatedFilePath)
	if fileInfo == nil {
		return nil, fmt.Errorf("Could not find file '%s.sh' in current directory\n.", commandName)
	}

	updatedFileContents, err := os.ReadFile(updatedFilePath)
	if err != nil {
		return nil, err
	}

	return updatedFileContents, nil
}
