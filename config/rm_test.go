package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveCommand(t *testing.T) {
	err := RemoveCommandByName("test")
	require.NoError(t, err)
}

func TestRemoveCommand2(t *testing.T) {
	err := RemoveCommandByName("test2")
	require.NoError(t, err)
}

func TestRemoveCommandById(t *testing.T) {
	err := RemoveCommandById("b6f5d20c-3fbe-4326-88c7-c3b4ef967c02")
	require.NoError(t, err)
}
