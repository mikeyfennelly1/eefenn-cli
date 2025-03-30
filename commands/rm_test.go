package commands

import (
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommand = subcommand.CreateSubCommand("test", "test-script.sh", "test command")

func TestRM(t *testing.T) {
	err := RemoveSubcommand(testSubcommand.Name)
	require.NoError(t, err)
}
