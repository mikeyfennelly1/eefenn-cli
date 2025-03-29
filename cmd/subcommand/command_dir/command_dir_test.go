package command_dir

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const TestCommandId = "b1e3130f-4a11-43e0-af39-aaffbcbfeccb"

func TestGetSubcommandDependenciesDirectory(t *testing.T) {
	actual := GetSubcommandDependenciesDirectory(TestCommandId)
	expected := EefennCLIRoot + "/" + TestCommandId + "/" + TestCommandId + ".dependencies"

	assert.Equal(t, actual, expected, "Strings should be equal")
}

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	actual := GetAbsoluteSubcommandDirname(TestCommandId)
	expected := EefennCLIRoot + "/" + TestCommandId

	assert.Equal(t, actual, expected, "Strings should be equal")
}

func TestCreateEmptySubcommandShellFile(t *testing.T) {
	_, err := CreateEmptySubcommandShellFile(TestCommandId)
	if err != nil {
		return
	}
}

func TestCreateSubcommandDirTree(t *testing.T) {
	err := CreateSubcommandDirTree(TestCommandId)
	if err != nil {
		return
	}
}

func TestRemoveCommandDirectoryRecursively(t *testing.T) {
	err := removeCommandDirectoryRecursively(TestCommandId)
	require.NoError(t, err)
}
