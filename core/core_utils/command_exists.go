package core_utils

func CommandExists(commandName string) bool {
	_, name, _ := GetCommand(commandName)
	if name != nil {
		return false
	} else {
		return false
	}
}
