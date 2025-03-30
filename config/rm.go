package config

func (config *Config) RemoveCommandByName(name string) error {
	var targetIndex int

	for index, scmd := range config.Subcommands {
		if scmd.Name == name {
			targetIndex = index
		}
	}

	config.Subcommands = append(config.Subcommands[:targetIndex], config.Subcommands[targetIndex+1:]...)

	config.Update()

	return nil
}

func RemoveCommandById(id string) error {
	config, err := GetCurrentConfig()
	if err != nil {
		return err
	}

	var targetIndex int

	for index, scmd := range config.Subcommands {
		if scmd.Hash[:8] == id {
			targetIndex = index
		}
	}

	config.Subcommands = append(config.Subcommands[:targetIndex], config.Subcommands[targetIndex+1:]...)

	config.Update()

	return nil
}
