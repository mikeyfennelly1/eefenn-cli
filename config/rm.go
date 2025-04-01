package config

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
