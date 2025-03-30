package config

// RemoveCommand
//
// remove a command from eefenn-cli.config.json by command name
func RemoveCommand(commandName string) error {
	configMap, err := getConfigMap()
	if err != nil {
		return err
	}

	delete(configMap, commandName)

	configByteSlice, err := getConfigByteSliceFromConfigMap(&configMap)
	if err != nil {
		return err
	}

	// write the updated slice of bytes to eefenn-cli.config.json
	err = writeToConfigFile(configByteSlice)
	if err != nil {
		return err
	}

	return nil
}
