package config

import (
	"github.com/eefenn/eefenn-cli/cmd-config"
)

// AddCommand
//
// Update /usr/lib/eefenn-cli/eefenn-cli.config.json with
// marshalled subcommand data.
func (c *Config) AddCommand(subcommand cmd_config.Command) error {
	c.Subcommands = append(c.Subcommands, subcommand)
	return nil
}
