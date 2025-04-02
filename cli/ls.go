package cli

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/core"
	"os"
	"text/tabwriter"
)

func LS() error {
	c, err := core.GetCore()
	if err != nil {
		return err
	}

	allCommands, err := c.GetALlCommands()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	// Print headers
	fmt.Fprintln(w, "NAME\tHAS ARGS")

	for _, command := range allCommands {
		hasArgs := "false"
		if len(command.Args) > 0 {
			hasArgs = "true"
		}
		fmt.Fprintf(w, "%s\t%s\n", command.Name, hasArgs)
	}

	// Flush the writer to ensure output
	w.Flush()

	return nil
}
