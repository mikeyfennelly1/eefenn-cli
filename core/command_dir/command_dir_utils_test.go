package command_dir

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSubCommandShellFileAbsPath(t *testing.T) {
	expected := "/usr/lib/eefenn-cli/test/test.sh"
	actual := getSubcommandShellFileAbsPath("test")
	assert.Equal(t, expected, actual)
}

func TestGetAbsoluteSubcommandDirname(t *testing.T) {
	expected := "/usr/lib/eefenn-cli/test"
	actual := getAbsoluteSubcommandDirname("test")
	assert.Equal(t, expected, actual)
}

func TestGetSubcommandDependenciesDirectory(t *testing.T) {
	expected := "/usr/lib/eefenn-cli/test/test.dependencies"
	actual := getSubcommandDependenciesDirectory("test")
	assert.Equal(t, expected, actual)
}
