package core

import (
	"github.com/eefenn/eefenn-cli/cmd"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCMD = cmd.Command{
	Name:        "test",
	Script:      "test.sh",
	Needs:       nil,
	Description: "test description for test-eefenn-cli.config.json",
	Args:        nil,
}

func TestEefennCLIDirectoryTree_CreateSubcommandDirTree(t *testing.T) {
	var edt eefennCLIDirectoryTree
	err := edt.CreateCMDDirTree(testCMD)
	require.NoError(t, err)
}
