package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCore_PrintSourceFiles(t *testing.T) {
	err := Run("cool-command", "/some/dir")
	require.NoError(t, err)
}
