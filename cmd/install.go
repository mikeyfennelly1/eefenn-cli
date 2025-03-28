// install.go
//
// 1. Builds the eefenn-cli binary
// 2. Appends the binary with an alias of eefenn to ~/.bashrc
// 3. Sources ~/.bashrc
// 4. Creates /usr/lib/eefenn directory
// 5. Creates eefenn-cli.config.json
// 6. Creates /usr/lib/eefenn/scripts
// 7. Installs contents of https://github.com/eefenn-cli/scripts to /usr/lib/eefenn/scripts

// ... Notifies user of exit status of the command

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const (
	UsrLibEefenn        = "/usr/lib/eefenn"
	EefennConfigJson    = "eefenn-cli.config.json"
	UsrLibEefennScripts = EefennConfigJson + "/scripts"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs efenn-cli.config.json, and latest eefenn cli shell scripts.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Installing...")
	},
}

// createUsrLibEefenn
//
// create the /usr/lib/eefenn direcory
func createUsrLibEefenn() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must be root to perform this operation")
	}

	err := os.Mkdir(UsrLibEefenn, 0666)
	if err != nil {
		return err
	}

	return nil
}

// createUsrLibEefenn
//
// create a blank configuration file for eefenn-cli
func createEefennCLIConfigJson() error {
	_, err := os.Create(EefennConfigJson)
	if err != nil {
		return err
	}

	return nil
}

// createUsrLibEefenn
//
// create the directory for scripts for this cli
func createScriptsDirectory() error {
	err := os.Mkdir(UsrLibEefennScripts, 0666)
	if err != nil {
		return err
	}

	return nil
}

// addRemoteToScriptsDir
//
// initialize the eefenn-cli scripts directory as a git repo
// add remote eefenn scripts repository as the remote for that repo
func addRemoteToScriptsDir() error {
	return nil
}
