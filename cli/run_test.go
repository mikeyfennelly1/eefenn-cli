package cli

import (
	"github.com/eefenn/eefenn-cli/core"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	pwd, _ := os.Getwd()
	err := core.Run("cool-command", pwd)
	require.NoError(t, err)
}
