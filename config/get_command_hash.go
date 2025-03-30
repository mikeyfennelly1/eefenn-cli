package config

import "fmt"

func (config *Config) GetCommandHash(name string) (*string, error) {
	for _, sc := range config.Subcommands {
		if sc.Name == name {
			return &sc.Hash, nil
		}
	}

	return nil, fmt.Errorf("Could not find hash for the command")
}
