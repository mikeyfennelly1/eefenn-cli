package subcommand

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type subcommandData struct {
	Hash        string `json:"command-hash"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type Config struct {
	Test subcommandData `json:"test"`
}

func (sc *subcommandData) Print() {
	var builder strings.Builder
	builder.WriteString("\t" + sc.Hash[:8])
	builder.WriteString("\t" + sc.Description)
	builder.WriteString("\t" + sc.Script + "\n")
	fmt.Printf(builder.String())
}

func ListCommands() error {
	eefennCliConfig, err := os.Open(EefennCLIConfig)
	if err != nil {
		return err
	}
	defer eefennCliConfig.Close()

	var config Config
	decoder := json.NewDecoder(eefennCliConfig)
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Printf("Error decoding %s: %v", EefennCLIConfig, err)
		return err
	}

	// Print the details of the "test" command
	config.Test.Print()

	return nil
}
