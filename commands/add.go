package commands

import (
	"github.com/eefenn/eefenn-cli/config"
	"github.com/google/uuid"
	"time"
)

func AddCommand(name string, description string, script string) error {
	config, err := config.ReadConfig()
	if err != nil {
		return err
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	subcommand := Subcommand{
		Name:        name,
		Script:      script,
		Hash:        uuid.String(),
		Description: description,
		DateCreated: time.DateTime,
	}

	config.AddCommand(subcommand)

	return nil
}
