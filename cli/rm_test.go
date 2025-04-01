package cli

import (
	"github.com/eefenn/eefenn-cli/yaml"
	"github.com/stretchr/testify/require"
	"testing"
)

var echoHello = yaml.Command{
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
