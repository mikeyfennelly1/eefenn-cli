package config

import (
	"github.com/eefenn/eefenn-cli/subcmd"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommand = subcmd.CreateSubCommand("test", "test-script.sh", "test command")
var testSubcommand2 = subcmd.CreateSubCommand("test2", "test-script.sh", "test command")

func TestRemoveCommand(t *testing.T) {
	err := RemoveCommand("test")
	require.NoError(t, err)

	err = RemoveCommand("test2")
	require.NoError(t, err)
}
