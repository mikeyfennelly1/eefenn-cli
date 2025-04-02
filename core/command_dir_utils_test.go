package core

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
	actual := getAbsoluteSubcommandDir("test")
	assert.Equal(t, expected, actual)
}

func TestGetSubcommandDependenciesDirectory(t *testing.T) {
	expected := "/usr/lib/eefenn-cli/test/test.dependencies"
	actual := getCMDDependenciesDir("test")
	assert.Equal(t, expected, actual)
}
