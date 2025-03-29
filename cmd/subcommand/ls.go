package subcommand

import (
	"encoding/json"
	"fmt"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type subcommandData struct {
	Hash        string `json:"command-hash"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type ConfigObject struct {
	Test subcommandData `json:"test"`
}

func GetConfigArray() (*ConfigObject, error) {
	eefennCliConfig, err := os.Open(EefennCLIConfig)
	if err != nil {
		return nil, err
	}
	defer eefennCliConfig.Close()

	var config ConfigObject
	decoder := json.NewDecoder(eefennCliConfig)
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Printf("Error decoding %s: %v", EefennCLIConfig, err)
		return nil, err
	}

	// Print the details of the "test" command
	return &config, nil
}
