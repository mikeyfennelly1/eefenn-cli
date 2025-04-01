package config

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddCommand(t *testing.T) {
	config, err := GetCurrentConfig()
	require.NoError(t, err)

	testCommandToAdd := cmd_config.Command{
		Name:        "test2",
		Script:      "test2.sh",
		Description: "Second test description",
	}

	config.AddCommand(testCommandToAdd)
	err = config.Update()

	require.NoError(t, err)
}
