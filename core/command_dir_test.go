package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEefennCLIDirectoryTree_CreateSubcommandDirTree(t *testing.T) {
	var edt eefennCLIDirectoryTree
	err := edt.CreateCMDDirTree(testCMD)
	require.NoError(t, err)
}
