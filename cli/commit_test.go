package cli

import (
	cmd_config "github.com/eefenn/eefenn-cli/cmd"
)

var TestCMD = cmd_config.Command{
	Name:   "test",
	Script: "test.sh",
}
