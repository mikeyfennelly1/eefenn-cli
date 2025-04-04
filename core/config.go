package core

import (
	"encoding/json"
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type ConfigInterface interface {
	// GetCommandByName
	//
	// Get a that exists in the config file command by the name of the command.
	//
	// Returns the index of the command, the Command struct for
	// the command, and error status.
	getCommandByName(name string) (*int, *cmd.Command, error)

	// AddCommand
	//
	// update /usr/lib/eefenn-cli/eefenn-cli.config.json with
	// marshalled subcommand data.
	addCMD(subcommand cmd.Command) error

	// GetCommandArgs
	//
	// Get the arguments to a command by commandName.
	//
	// Returns command's arguments and error status.
	getCommandArgs(commandName string) ([]cmd.Arg, error)

	// RemoveCommandByName
	//
	// remove a command using the command name as a parameter.
	removeCommandByName(name string) error
}

type config struct {
	RemoteRepoURL string        `json:"remoteRepoURL"`
	Commands      []cmd.Command `json:"commands"`
}

// GetCommandArgs
//
// Get the arguments to a command by commandName.
//
// Returns command's arguments and error status.
func (c *config) getCommandArgs(commandName string) ([]cmd.Arg, error) {
	//TODO implement me
	panic("implement me")
}

// getCurrentConfig
//
// Get a config object for the current state of the config file.
//
// Returns config object, and error status.
func getCurrentConfig() (config, error) {
	var config config

	// Open the JSON file
	file, err := os.Open(EefennCLIConfig)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Decode the JSON data into the struct
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// GetCommandByName
//
// Get a that exists in the config file command by the name of the command.
//
// Returns the index of the command, the Command struct for
// the command, and error status.
func (c *config) getCommandByName(name string) (*int, *cmd.Command, error) {
	var p_cmd *cmd.Command
	var p_commandIndex *int

	// find the index of the item whose Name matches the parameter 'name'
	for index, sc := range c.Commands {
		if sc.Name == name {
			p_commandIndex = &index
			p_cmd = &sc
		}
	}

	if p_cmd == nil {
		return nil, nil, fmt.Errorf("Could not find command with the name: %s\n", name)
	} else {
		return p_commandIndex, p_cmd, nil
	}
}

// update
//
// Write the contents of a config object to the config file.
//
// Returns error status.
func (c *config) update() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You must be root to update commands.\n")
	}
	jsonData, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(EefennCLIConfig, jsonData, 0666)
	if err != nil {
		return err
	}

	return nil
}

// AddCommand
//
// update /usr/lib/eefenn-cli/eefenn-cli.config.json with
// marshalled subcommand data.
func (c *config) addCMD(cmd cmd.Command) error {
	c.Commands = append(c.Commands, cmd)
	err := c.update()
	if err != nil {
		return err
	}

	return nil
}

// RemoveCommandByName
//
// remove a command using the command name as a parameter.
func (config *config) removeCommandByName(name string) error {
	var targetIndex int

	for index, existingCommand := range config.Commands {
		if existingCommand.Name == name {
			targetIndex = index
		}
	}

	config.Commands = append(config.Commands[:targetIndex], config.Commands[targetIndex+1:]...)

	err := config.update()
	if err != nil {
		return err
	}

	return nil
}
