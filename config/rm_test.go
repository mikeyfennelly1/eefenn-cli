package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveCommandByName(t *testing.T) {
	currentConfig, err := GetCurrentConfig()
	err = currentConfig.RemoveCommandByName("test")
	require.NoError(t, err)
}

func TestRemoveCommandById(t *testing.T) {
	err := RemoveCommandById("b6f5d20c-3fbe-4326-88c7-c3b4ef967c02")
	require.NoError(t, err)
}
