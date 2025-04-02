package core

func cmdExists(currentCore Core, commandName string) bool {
	for _, cmd := range currentCore.GetCommands() {
		if cmd.Name == commandName {
			return true
		}
	}

	return false
}
