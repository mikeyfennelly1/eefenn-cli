package add_subcommand

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

var testSubcommandDependencies = []string{"dep1.txt", "dep2.txt"}
var testSubcommand = CreateSubCommand("test", "test-script.sh", testSubcommandDependencies, "test command")

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	absoluteSubcommandDirname := testSubcommand.getAbsoluteSubcommandDirname()
	fmt.Println(absoluteSubcommandDirname)
}

func TestUpdateConfig(t *testing.T) {
	err := testSubcommand.updateConfig()
	require.NoError(t, err)
}

func TestCreateSubcommandDirTree(t *testing.T) {
	err := testSubcommand.createSubcommandDirTree()
	require.NoError(t, err)
}

func TestUpdateConfigJSON(t *testing.T) {
	err := testSubcommand.updateConfig()
	require.NoError(t, err)
}
