package config

import (
	"encoding/json"
	"fmt"
	"github.com/eefenn/eefenn-cli/yaml"
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
	GetCommandByName(name string) (*int, *yaml.Command, error)

	// Update
	//
	// Write the contents of a Config object to the config file.
	//
	// Returns error status.
	Update() error

	// AddCommand
	//
	// Update /usr/lib/eefenn-cli/eefenn-cli.config.json with
	// marshalled subcommand data.
	AddCommand(subcommand yaml.Command) error

	// GetCommandArgs
	//
	// Get the arguments to a command by commandName.
	//
	// Returns command's arguments and error status.
	GetCommandArgs(commandName string) ([]yaml.Arg, error)

	// RemoveCommandByName
	//
	// remove a command using the command name as a parameter.
	RemoveCommandByName(name string) error
}

type Config struct {
	RemoteRepoURL string         `json:"remoteRepoURL"`
	Commands      []yaml.Command `json:"subcommands"`
}

// GetCommandArgs
//
// Get the arguments to a command by commandName.
//
// Returns command's arguments and error status.
func (c *Config) GetCommandArgs(commandName string) ([]yaml.Arg, error) {
	//TODO implement me
	panic("implement me")
}

// GetCurrentConfig
//
// Get a Config object for the current state of the config file.
//
// Returns Config object, and error status.
func GetCurrentConfig() (Config, error) {
	var config Config

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
func (c *Config) GetCommandByName(name string) (*int, *yaml.Command, error) {
	var p_cmd *yaml.Command
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

// Update
//
// Write the contents of a Config object to the config file.
//
// Returns error status.
func (c *Config) Update() error {
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
// Update /usr/lib/eefenn-cli/eefenn-cli.config.json with
// marshalled subcommand data.
func (c *Config) AddCommand(subcommand yaml.Command) error {
	c.Commands = append(c.Commands, subcommand)
	return nil
}

// RemoveCommandByName
//
// remove a command using the command name as a parameter.
func (config *Config) RemoveCommandByName(name string) error {
	var targetIndex int

	for index, scmd := range config.Commands {
		if scmd.Name == name {
			targetIndex = index
		}
	}

	config.Commands = append(config.Commands[:targetIndex], config.Commands[targetIndex+1:]...)

	config.Update()

	return nil
}
