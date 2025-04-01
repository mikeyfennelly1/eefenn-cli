package cli

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/require"
	"testing"
)

var echoHello = cmd_config.Command{
	Name:        "echo-hello",
	Script:      "echo-hello.sh",
	Needs:       nil,
	Description: "",
	Args:        nil,
}

func TestRM(t *testing.T) {
	err := RemoveSubcommand(echoHello.Name)
	require.NoError(t, err)
}
