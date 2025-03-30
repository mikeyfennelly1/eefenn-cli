package command_dir

import (
	"io"
	"os"
)

func CopyShellScript(sourceScript string, commandHash string) error {
	sourceFile, err := os.OpenFile(sourceScript, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := CreateEmptySubcommandShellFile(commandHash)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync() // Ensure all writes are flushed to disk
}
