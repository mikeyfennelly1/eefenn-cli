package config

import (
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCMD = cmd.Command{
	Name:        "test",
	Script:      "test.sh",
	Needs:       nil,
	Description: "test description for test-eefenn-cli.config.json",
	Args:        nil,
}

func TestGetCurrentConfig(t *testing.T) {
	_, err := GetCurrentConfig()
	require.NoError(t, err)
}

func TestConfig_AddCommand(t *testing.T) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return
	}

	err = currentConfig.AddCommand(testCMD)
	require.NoError(t, err)
}

func TestConfig_RemoveCommandByName(t *testing.T) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return
	}

	err = currentConfig.RemoveCommandByName(testCMD.Name)
	require.NoError(t, err)
}
