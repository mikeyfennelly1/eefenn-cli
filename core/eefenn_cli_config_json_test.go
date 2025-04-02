package core

import (
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCMD = cmd.Command{
	Name:         "test",
	Script:       "test.sh",
	Dependencies: nil,
	Description:  "test description for test-eefenn-cli.config.json",
	Args:         nil,
}

func TestGetCurrentConfig(t *testing.T) {
	_, err := getCurrentConfig()
	require.NoError(t, err)
}

func TestConfig_AddCommand(t *testing.T) {
	currentConfig, err := getCurrentConfig()
	require.NoError(t, err)

	err = currentConfig.addCMD(testCMD)
	require.NoError(t, err)
}

func TestConfig_GetCommandByName(t *testing.T) {
	currentConfig, err := getCurrentConfig()
	require.NoError(t, err)
	_, result, err := currentConfig.getCommandByName(testCMD.Name)
	require.NoError(t, err)
	assert.Equal(t, testCMD, *result)
}

func TestConfig_RemoveCommandByName(t *testing.T) {
	currentConfig, err := getCurrentConfig()
	if err != nil {
		return
	}

	err = currentConfig.removeCommandByName(testCMD.Name)
	require.NoError(t, err)
}
