package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCMDExists(t *testing.T) {
	exists := cmdExists("test-command")
	assert.False(t, exists)
	fmt.Printf("%v\n", exists)
}
