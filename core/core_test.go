package core

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCore_PrintSourceFiles(t *testing.T) {
	_, err := getImgFilesToRunFilesMap("cool-command", "/some/dir")
	require.NoError(t, err)
}

func TestCore_RunCommand(t *testing.T) {
	pwd, _ := os.Getwd()
	err := Run("cool-command", pwd)
	require.NoError(t, err)
}
