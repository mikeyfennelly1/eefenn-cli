package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCore_PrintSourceFiles(t *testing.T) {
	err := PrintSourceFiles("cool-command")
	require.NoError(t, err)
}
