package subcommand

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListCommands(t *testing.T) {
	err := ListCommands()
	require.NoError(t, err)
}
