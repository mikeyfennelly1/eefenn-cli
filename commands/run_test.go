package commands

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	output, err := Run("echo-hello", nil)
	fmt.Printf(string(output) + "\n")
	require.NoError(t, err)
}
