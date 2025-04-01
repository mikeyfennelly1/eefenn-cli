package cli

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCat(t *testing.T) {
	err := Cat("test2")
	require.NoError(t, err)
}
