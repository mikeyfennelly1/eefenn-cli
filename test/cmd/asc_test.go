package test

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd/asc"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommandDependencies = []string{"dep1.txt", "dep2.txt"}
var testSubcommand = asc.CreateSubCommand("test", "test-script.sh", testSubcommandDependencies, "test command")

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	absoluteSubcommandDirname := testSubcommand.GetAbsoluteSubcommandDirname()
	fmt.Println(absoluteSubcommandDirname)
}

func TestCreateSubCommandConfigEntry(t *testing.T) {
	_ = testSubcommand.CreateSubCommandConfigEntry()
}

func TestCreateSubcommandDirTree(t *testing.T) {
	err := testSubcommand.CreateSubcommandDirTree()
	require.NoError(t, err)
}

func TestUpdateConfigJSON(t *testing.T) {
	err := testSubcommand.UpdateConfigJSON()
	require.NoError(t, err)
}
