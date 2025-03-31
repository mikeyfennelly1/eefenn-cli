package config

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSubcommands = []cmd_config.Subcommand{
	{
		Name:        "test",
		Script:      "test.sh",
		Hash:        "b6f5d20c-3fbe-4326-88c7-c3b4ef967c02",
		Description: "test description for test-eefenn-cli.config.json",
		DateCreated: "2025-03-30T14:25:36Z",
	},
}

func TestReadConfig(t *testing.T) {
	config, err := GetCurrentConfig()
	if err != nil {
		return
	}

	assert.Equal(t, testSubcommands, config.Subcommands)
}
