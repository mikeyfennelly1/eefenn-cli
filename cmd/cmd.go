package cmd

import (
	"fmt"
	"os"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

type Command struct {
	// alias of the script
	Name string `yaml:"name" json:"name"`

	// the script which the command is an alias for
	Script string `yaml:"script" json:"script"`

	// array of filepaths that the script needs to run
	Dependencies []string `yaml:"dependencies" json:"dependencies"`

	// description for what the script does
	Description string `yaml:"description" json:"description"`

	// the arguments to the command
	Args []Arg ` yaml:"args" json:"args"`
}

type Arg struct {
	Name        string `json:"name" json:"name"`
	Type        string `json:"type" json:"type"`
	Description string `json:"description" json:"description"`
}

func (c *Command) GetCmdFilePaths() []string {
	var filePaths []string

	for _, file := range c.Dependencies {
		filePaths = append(filePaths, file)
	}
	filePaths = append(filePaths, c.Script)

	return filePaths
}

func (c *Command) GetCmdImgDirPath() string {
	return fmt.Sprintf("%s/%s", EefennCLIRoot, c.Name)
}

func (c *Command) GetCommandScriptPathInRunDir(runDir string) string {
	return runDir + c.Script
}

func (c *Command) CreateCommandImageDir() error {
	imgDir := c.GetCmdImgDirPath()

	err := os.MkdirAll(imgDir, 0775)
	if err != nil {
		return err
	}

	return nil
}
