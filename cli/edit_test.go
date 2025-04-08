package cli

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEdit(t *testing.T) {
	err := Edit("cool-command")
	require.NoError(t, err)
}
