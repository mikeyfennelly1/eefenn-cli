package command_dir

import (
	cmd_config "github.com/eefenn/eefenn-cli/cmd-config"
	"io"
	"os"
)

func CopyShellScript(cmd cmd_config.Command) error {
	sourceFile, err := os.OpenFile(cmd.Script, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := CreateEmptySubcommandShellFile(cmd)
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
