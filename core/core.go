package core

import (
	"github.com/eefenn/eefenn-cli/core/command_dir"
	"github.com/eefenn/eefenn-cli/core/config"
	cmd_config "github.com/eefenn/eefenn-cli/yaml"
)

type CoreInterface interface {
	Commit(command cmd_config.Command)

	List()
}

type Core struct {
	Config        config.Config
	DirectoryTree command_dir.EefennCLIDirectoryTree
}

// Commit
//
// Add/'commit' a command to core.
func (c *Core) Commit(command cmd_config.Command) {

}

func (c *Core) GetCommands() []cmd_config.Command {
	return nil
}
