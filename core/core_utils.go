package core

import "fmt"

func CmdExists(currentCore CoreInterface, commandName string) bool {
	for _, cmd := range currentCore.GetALlCommands() {
		if cmd.Name == commandName {
			fmt.Printf("This is true, '%s' is the same as '%s'", cmd.Name, commandName)
			return true
		}
	}

	return false
}
