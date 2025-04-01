package config

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSubcommands = []cmd_config.Command{
	{
		Name:        "test",
		Script:      "test.sh",
		Description: "test description for test-eefenn-cli.config.json",
	},
}

func TestReadConfig(t *testing.T) {
	config, err := GetCurrentConfig()
	if err != nil {
		return
	}

	assert.Equal(t, testSubcommands, config.Commands)
}
