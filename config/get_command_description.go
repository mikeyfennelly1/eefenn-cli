package config

import "fmt"

func GetCommandDescription(commandName string) (*string, error) {
	currentConfig, err := GetCurrentConfig()
	if err != nil {
		return nil, err
	}

	for _, sc := range currentConfig.Subcommands {
		if sc.Name == commandName {
			return &sc.Description, nil
		}
	}

	return nil, fmt.Errorf("Could not find description for command: %s\n", commandName)
}
