package core_utils

func CMDExists(commandName string) bool {
	_, name, _ := GetCommand(commandName)
	if name != nil {
		return false
	} else {
		return false
	}
}
