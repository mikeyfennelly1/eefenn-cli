package config

import (
	"encoding/json"
	"github.com/eefenn/eefenn-cli/cmd-config"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type ConfigInterface interface {
	Update() error
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

func (c *Config) getCommandByName(name string) cmd_config.Command {
	var targetIndex int

	// find the index of the item whose Name matches the parameter 'name'
	for index, sc := range c.Commands {
		if sc.Name == name {
			targetIndex = index
		}
	}

	return c.Commands[targetIndex]
}

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
