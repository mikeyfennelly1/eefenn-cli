package core

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnMarshalCommandFromYamlContents_NameScriptDescription(t *testing.T) {
	input := "name: test-CommandImage\nscript: test-eefenn.sh\ndescription: This is a test CommandImage."
	expected := &CommandInterface{
		Name:        "test-CommandImage",
		Script:      "test-eefenn.sh",
		Description: "This is a test CommandImage.",
	}
	actual, _ := unMarshalCommandFromYamlContents([]byte(input))

	assert.Equal(t, expected, actual)
}

func TestUnMarshalCommandFromYamlContents_NameScriptDescriptionArgs(t *testing.T) {
	input := "name: test-CommandImage\nscript: test-eefenn.sh\ndescription: This is a test CommandImage.\nneeds:\n  - ./file1\n  - ./file2\nargs:\n  - name: arg1\n    type: string\n    description: A test argument"
	expected := &CommandInterface{
		Name:         "test-CommandImage",
		Script:       "test-eefenn.sh",
		Description:  "This is a test CommandImage.",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
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
	validCMD := &CommandInterface{
		Name:         "test-CommandImage",
		Script:       "test-eefenn.sh",
		Description:  "This is a test CommandImage.",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual := validateCMDSyntax(validCMD)

	require.NoError(t, actual)

	// test a CommandImage where there is no name
	noNameCMD := &CommandInterface{
		Name:         "",
		Script:       "test-eefenn.sh",
		Description:  "This is a test CommandImage.",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(noNameCMD)

	require.Error(t, actual)

	// test a CommandImage where the description is empty string
	noDescriptionCMD := &CommandInterface{
		Name:         "test-cmd",
		Script:       "test-eefenn.sh",
		Description:  "",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
				Name:        "arg1",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(noDescriptionCMD)

	require.Error(t, actual)

	// test a CommandImage where the arg type is invalid
	invalidArgCMD := &CommandInterface{
		Name:         "test-cmd",
		Script:       "test-eefenn.sh",
		Description:  "description",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
				Name:        "arg1",
				Type:        "invalid",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(invalidArgCMD)

	require.Error(t, actual)

	// test a CommandImage where the arg type is invalid
	emptyArgNameCMD := &CommandInterface{
		Name:         "test-cmd",
		Script:       "test-eefenn.sh",
		Description:  "description",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
				Name:        "",
				Type:        "string",
				Description: "A test argument",
			},
		},
	}
	actual = validateCMDSyntax(emptyArgNameCMD)

	require.Error(t, actual)

	// test a CommandImage where the arg type is invalid
	emptyArgDescriptionCMD := &CommandInterface{
		Name:         "test-cmd",
		Script:       "test-eefenn.sh",
		Description:  "description",
		Dependencies: []string{"./file1", "./file2"},
		Args: []Arg{
			{
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
