package subcmd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommand = CreateSubCommand("test", "test-script.sh", "test command")

func TestAddSubCommand(t *testing.T) {
	err := testSubcommand.AddSubCommand()
	require.NoError(t, err)
}

func TestRemoveSubcommand(t *testing.T) {
	err := RemoveSubcommand(testSubcommand.Hash.String(), testSubcommand.Name)
	require.NoError(t, err)
}
