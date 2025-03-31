package cli

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCat(t *testing.T) {
	err := Cat("echo-hello")
	require.NoError(t, err)
}
