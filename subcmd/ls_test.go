package subcmd

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListCommands(t *testing.T) {
	err := LS()
	require.NoError(t, err)
}
