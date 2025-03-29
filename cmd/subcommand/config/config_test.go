package config

import (
	"github.com/eefenn/eefenn-cli/cmd/subcommand"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommand = subcommand.CreateSubCommand("test", "test-script.sh", "test command")
var testSubcommand2 = subcommand.CreateSubCommand("test2", "test-script.sh", "test command")

func TestAddCommand(t *testing.T) {
	err := AddCommand(&testSubcommand)
	require.NoError(t, err)

	err = AddCommand(&testSubcommand2)
	require.NoError(t, err)
}

func TestRemoveCommand(t *testing.T) {
	err := RemoveCommand("test")
	require.NoError(t, err)

	err = RemoveCommand("test2")
	require.NoError(t, err)
}
