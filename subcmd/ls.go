package subcmd

import (
	"encoding/json"
	"fmt"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type subcommandData struct {
	ID          string `json:"command-hash"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type Config struct {
	Test subcommandData `json:"test"`
}

type CommandPrintFormat struct {
	Name string
	Hash string
}

func printHeaders() {
	headers := []string{"ID", "NAME"}
	fmt.Printf("%-10s %-20s\n", headers[0], headers[1])

	return
}

func (cpf *CommandPrintFormat) printCommandLine() {
	headers := []string{cpf.Hash, cpf.Name}
	fmt.Printf("%-10s %-20s\n", headers[0], headers[1])

	return

}

func (sc *subcommandData) getPrintFormat() (*CommandPrintFormat, error) {
	if len(sc.ID) <= 7 {
		return nil, fmt.Errorf("ID is not long enough to print\n")
	}

	printFormat := &CommandPrintFormat{
		Name: "placeholder",
		Hash: sc.ID[:8],
	}
	return printFormat, nil
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
	configAsPrintFormat, err := config.Test.getPrintFormat()
	if err != nil {
		return err
	}

	printHeaders()

	configAsPrintFormat.printCommandLine()

	return nil
}
