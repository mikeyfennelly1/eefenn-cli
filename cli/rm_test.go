package cli

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommand = cmd_config.CreateSubCommand("test", "test-script.sh", "test command")

func TestRM(t *testing.T) {
	err := RemoveSubcommand(testSubcommand.Name)
	require.NoError(t, err)
}
