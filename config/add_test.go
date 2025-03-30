package config

import (
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddCommand(t *testing.T) {
	config, err := ReadConfig()
	require.NoError(t, err)

	testCommandToAdd := subcommand.Subcommand{
		Name:        "test2",
		Script:      "test2.sh",
		Hash:        "b6f5d20c-3fbe-4326-88c7-c3b4ef967c02",
		Description: "Second test description",
	}

	config.AddCommand(testCommandToAdd)
	err = config.Update()

	require.NoError(t, err)
}
