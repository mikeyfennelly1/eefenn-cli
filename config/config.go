package config

import (
	"encoding/json"
	"github.com/eefenn/eefenn-cli/subcommand"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type Config struct {
	RemoteRepoURL string
	Subcommands   []subcommand.Subcommand
}

// writeToConfigFile
//
// write the contents of a byte array to the eefenn-cli.config.json
func writeToConfigFile(updatedConfig []byte) error {
	// Write directly to file without re-marshaling
	err := os.WriteFile(EefennCLIConfig, updatedConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig() (Config, error) {
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

func (c *Config) getSubCommandByName(name string) subcommand.Subcommand {
	var targetIndex int

	// find the index of the item whose Name matches the parameter 'name'
	for index, sc := range c.Subcommands {
		if sc.Name == name {
			targetIndex = index
		}
	}

	return c.Subcommands[targetIndex]
}
