// ls.go
//
// list subcommands in this CLI

package ls

import (
	"encoding/json"
	"fmt"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli.config.json"

func ListCommands() error {
	file, err := os.Open(EefennCLIConfig)
	if err != nil {
		return fmt.Errorf("Could not open %s\n", EefennCLIConfig)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode()

	return nil
}
