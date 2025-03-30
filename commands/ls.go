package commands

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/config"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

func LS() error {
	config, err := config.ReadConfig()
	if err != nil {
		return err
	}

	printHeaders()

	for _, sc := range config.Subcommands {
		sc.List()
	}

	return nil
}

func printHeaders() {
	headers := []string{"ID", "NAME", "DESCRIPTION"}
	fmt.Printf("%-10s %-10s %-20s\n", headers[0], headers[1], headers[2])

	return
}
