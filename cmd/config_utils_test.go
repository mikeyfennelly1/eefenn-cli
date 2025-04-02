package cmd

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnMarshalCommandFromYamlContents_NameScriptDescription(t *testing.T) {
	input := "name: test-command\nscript: test-eefenn.sh\ndescription: This is a test command."
	expected := &Command{
		Name:        "test-command",
		Script:      "test-eefenn.sh",
		Description: "This is a test command.",
	}
	actual, _ := unMarshalCommandFromYamlContents([]byte(input))

	assert.Equal(t, expected, actual)
}

func TestUnMarshalCommandFromYamlContents_NameScriptDescriptionArgs(t *testing.T) {
	input := "name: test-command\nscript: test-eefenn.sh\ndescription: This is a test command.\nneeds:\n  - ./file1\n  - ./file2\nargs:\n  - name: arg1\n    type: string\n    description: A test argument"
	expected := &Command{
		Name:        "test-command",
		Script:      "test-eefenn.sh",
		Description: "This is a test command.",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual, _ := unMarshalCommandFromYamlContents([]byte(input))

	assert.Equal(t, expected, actual)
}

func TestValidateCMDSyntax(t *testing.T) {
	validCMD := &Command{
		Name:        "test-command",
		Script:      "test-eefenn.sh",
		Description: "This is a test command.",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual := validateCMDSyntax(validCMD)

	require.NoError(t, actual)

	// test a command where there is no name
	noNameCMD := &Command{
		Name:        "",
		Script:      "test-eefenn.sh",
		Description: "This is a test command.",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(noNameCMD)

	require.Error(t, actual)

	// test a command where the description is empty string
	noDescriptionCMD := &Command{
		Name:        "test-cmd",
		Script:      "test-eefenn.sh",
		Description: "",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(noDescriptionCMD)

	require.Error(t, actual)

	// test a command where the arg type is invalid
	invalidArgCMD := &Command{
		Name:        "test-cmd",
		Script:      "test-eefenn.sh",
		Description: "description",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "arg1",
				Type:        "invalid",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(invalidArgCMD)

	require.Error(t, actual)

	// test a command where the arg type is invalid
	emptyArgNameCMD := &Command{
		Name:        "test-cmd",
		Script:      "test-eefenn.sh",
		Description: "description",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(emptyArgNameCMD)

	require.Error(t, actual)

	// test a command where the arg type is invalid
	emptyArgDescriptionCMD := &Command{
		Name:        "test-cmd",
		Script:      "test-eefenn.sh",
		Description: "description",
		Needs:       []string{"./file1", "./file2"},
		Args: []Arg{
			Arg{
				Name:        "name",
				Type:        "string",
				Description: "",
			},
		},
	}
	actual = validateCMDSyntax(emptyArgDescriptionCMD)

	require.Error(t, actual)

}

func TestGetCMDFromPWD(t *testing.T) {
	cmd, err := GetCMDFromPWD()
	require.NoError(t, err)
	fmt.Printf(cmd.Name)
}
