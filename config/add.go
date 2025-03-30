package config

import "github.com/eefenn/eefenn-cli/subcommand"

// AddCommand
//
// Update /usr/lib/eefenn-cli/eefenn-cli.config.json with
// marshalled subcommand data.
func (c *Config) AddCommand(subcommand subcommand.Subcommand) {
	c.Subcommands = append(c.Subcommands, subcommand)
	return
}

func (c *Config) getSubCommandByID(hash string) {
	var targetIndex int

	// find the index of the item whose hash matches the parameter 'hash'
	for index, sc := range c.Subcommands {
		if sc.Hash[:8] == hash {
			targetIndex = index
		}
	}

	// update the subcommands slice to exclude the subcommand at targetIndex
	c.Subcommands = append(c.Subcommands[:targetIndex], c.Subcommands[targetIndex+1:]...)
}
