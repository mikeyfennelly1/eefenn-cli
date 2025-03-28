package asc

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommandDependencies = []string{"dep1.txt", "dep2.txt"}
var testSubcommand = createSubCommand("test", "test-script.sh", testSubcommandDependencies, "test command")

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	absoluteSubcommandDirname := testSubcommand.getAbsoluteSubcommandDirname()
	fmt.Println(absoluteSubcommandDirname)
}

func TestCreateSubCommandConfigEntry(t *testing.T) {
	_ = testSubcommand.createSubCommandConfigEntry()
}

func TestCreateSubcommandDirTree(t *testing.T) {
	err := testSubcommand.createSubcommandDirTree()
	require.NoError(t, err)
}

func TestUpdateConfigJSON(t *testing.T) {
	err := testSubcommand.updateConfigJSON()
	require.NoError(t, err)
}
