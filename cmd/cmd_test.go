package cmd

import (
	"testing"
)

func TestGetCommandFromYml(t *testing.T) {
	ParseCommandFromYaml("config.yaml")
}
