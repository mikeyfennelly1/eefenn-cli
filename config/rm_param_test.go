package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRmParam(t *testing.T) {
	err := RmParam("echo-hello", "test-param")
	require.NoError(t, err)
}
