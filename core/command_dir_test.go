package command_dir

import (
	"fmt"
	cmd_config "github.com/eefenn/eefenn-cli/cmd-config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var cmd = cmd_config.Command{
	Name:        "create-next-component",
	Script:      fmt.Sprintf("%s/9c87b04f-71e4-488b-8b91-c1f8a9f56856/%s", EefennCLIRoot, "create-next-component.sh"),
	Needs:       nil,
	Description: "clean-next-project",
	Args:        nil,
}

func TestGetSubcommandDependenciesDirectory(t *testing.T) {
	actual := GetSubcommandDependenciesDirectory(cmd)
	expected := EefennCLIRoot + "/" + cmd.Name + "/" + cmd.Name + ".dependencies"

	assert.Equal(t, actual, expected, "Strings should be equal")
}

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	actual := GetAbsoluteSubcommandDirname(cmd.Name)
	expected := EefennCLIRoot + "/" + cmd.Name

	assert.Equal(t, actual, expected, "Strings should be equal")
}

func TestCreateEmptySubcommandShellFile(t *testing.T) {
	_, err := CreateEmptySubcommandShellFile(cmd)
	if err != nil {
		return
	}
}

func TestCreateSubcommandDirTree(t *testing.T) {
	err := CreateSubcommandDirTree(cmd)
	if err != nil {
		return
	}
}

func TestRemoveCommandDirectoryRecursively(t *testing.T) {
	err := RemoveCommandDirectoryRecursively(cmd.Name)
	require.NoError(t, err)
}
