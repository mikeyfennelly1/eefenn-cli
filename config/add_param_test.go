package config

import (
	"github.com/eefenn/eefenn-cli/subcommand"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddParam(t *testing.T) {
	paramToAdd := subcommand.Parameter{
		Name:        "test-param",
		Description: "A test parameter",
	}
	err := AddParam("echo-hello", paramToAdd)
	require.NoError(t, err)
}
