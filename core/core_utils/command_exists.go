package core_utils

import (
	"strings"
)

func CMDExists(commandName string) bool {
	var cmdExists bool
	_, cmd, _ := GetCommand(commandName)

	if cmd != nil {
		cmdExists = strings.Compare(cmd.Name, commandName) == 0
	} else {
		cmdExists = false
	}
	return cmdExists
}
