package cli

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEdit(t *testing.T) {
	err := Edit("create-next-component")
	require.NoError(t, err)
}
