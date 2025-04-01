package cli

import (
	cmd_config "github.com/eefenn/eefenn-cli/cmd-config"
	"testing"
)

var TestCMD = cmd_config.Command{
	Name:   "test",
	Script: "test.sh",
}

func TestAdd(t *testing.T) {
	Add()
}
