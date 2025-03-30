package commands

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	output, err := Run("echo-hello")
	fmt.Printf(string(output) + "\n")
	require.NoError(t, err)
}
