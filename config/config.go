package config

import (
	"encoding/json"
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd-config"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type ConfigInterface interface {
	GetCommandByName(name string) (*int, *cmd_config.Command, error)

	Update() error

	AddCommand(subcommand cmd_config.Command)

	GetCommandArgs(commandName string) ([]cmd_config.Arg, error)

	RemoveCommandByName(name string) error
}

type Config struct {
	RemoteRepoURL string               `json:"remoteRepoURL"`
	Commands      []cmd_config.Command `json:"subcommands"`
}

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
func (c *Config) GetCommandByName(name string) (*int, *cmd_config.Command, error) {
	var p_cmd *cmd_config.Command
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
func (c *Config) AddCommand(subcommand cmd_config.Command) error {
	c.Commands = append(c.Commands, subcommand)
	return nil
}

// GetCommandArgs
//
// Get the arguments of a command by command name
func GetCommandArgs(commandName string) ([]cmd_config.Arg, error) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	for _, sc := range currentConfig.Commands {
		if sc.Name == commandName {
			return sc.Args, nil
		}
	}

	return nil, fmt.Errorf("Could not find parameters for the command: '%s'\n", commandName)
}

// RemoveCommandByName
//
// remove a command using the command name as a parameter
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
