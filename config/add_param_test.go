package config

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddParam(t *testing.T) {
	paramToAdd := cmd_config.Parameter{
		Name:        "test-param",
		Description: "A test parameter",
	}
	err := AddParam("echo-hello", paramToAdd)
	require.NoError(t, err)
}
