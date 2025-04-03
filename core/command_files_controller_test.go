package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEefennCLIDirectoryTree_CreateSubcommandDirTree(t *testing.T) {
	var edt cmdFilesController
	err := edt.createCMDDir(testCMD)
	require.NoError(t, err)
}
