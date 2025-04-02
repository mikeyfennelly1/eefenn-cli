package core_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCMDExists(t *testing.T) {
	exists := CMDExists("test-command")
	assert.False(t, exists)
	fmt.Printf("%v\n", exists)
}
