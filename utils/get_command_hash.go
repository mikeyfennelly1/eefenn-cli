package utils

// GetCommandHash
//
// get the hash of a command by commandName
func GetCommandHash(commandName string) (*string, error) {
	_, command, err := GetCommand(commandName)
	if err != nil {
		return nil, err
	}

	return &command.Hash, nil
}
