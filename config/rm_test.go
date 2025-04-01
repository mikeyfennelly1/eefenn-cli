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
