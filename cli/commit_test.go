package cli

import (
	cmd_config "github.com/eefenn/eefenn-cli/cmd"
)

var TestCMD = cmd_config.CommandInterface{
	Name:   "test",
	Script: "test.sh",
}
