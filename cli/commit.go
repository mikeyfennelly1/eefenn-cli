package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/eefenn/eefenn-cli/core"
	"os"
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

	// copy the command's config.yaml to pwd
	configYamlInImageDir := fmt.Sprintf("%s/%s/config.yaml", core.EefennCLIRoot, pCMD.Name)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	configYamlInPWD := fmt.Sprintf("%s/config.yaml", pwd)
	err = core.CopyFile(configYamlInPWD, configYamlInImageDir)
	if err != nil {
		return err
	}

	fmt.Printf("Added command: %s\n", (*pCMD).Name)

	return nil
}
