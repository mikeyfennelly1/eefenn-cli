// core.go
//
// core is essentially the management system for commands, and their files.
//
// Author: Mikey Fennelly

package core

import (
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func GetCore() (CoreInterface, error) {
	config, err := getCurrentConfig()
	if err != nil {
		return nil, err
	}

	var edt cmdFilesController

	current_core := core{
		config:        config,
		cmdController: edt,
	}

	return &current_core, nil
}

type CoreInterface interface {
	// Commit
	//
	// Add/'commit' a command to core.
	Commit(command cmd.CommandInterface) error

	// GetCommandByName
	//
	// Get a command object for a command, using the name of the command as
	// a parameter.
	GetCommandByName(commandName string) (*cmd.CommandInterface, error)

	// GetALlCommands
	//
	// Get all commands in the current core state.
	GetALlCommands() ([]cmd.CommandInterface, error)

	// RemoveCommandByName
	//
	// Remove a command, specifying which command by name of the command.
	RemoveCommandByName(commandName string) error

	// CreateRunnableCommandInstance
	//
	// Create a command from a command image.
	CreateRunnableCommandInstance(targetDirectory absolutePath, imageInterface commandImageInterface) RunnableCommandInterface
}

type core struct {
	config        config
	cmdController cmdFilesController
}

func (c *core) GetCommandByName(commandName string) (*cmd.CommandInterface, error) {
	for _, command := range c.config.Commands {
		if command.Name == commandName {
			return &command, nil
		}
	}

	return nil, fmt.Errorf("command does not exist: %s", commandName)
}

// GetALlCommands
// Gets all commands in config. If there are no commands,
// will return an error
func (c *core) GetALlCommands() ([]cmd.CommandInterface, error) {
	if len(c.config.Commands) == 0 {
		return nil, fmt.Errorf("there are no commands")
	}
	return c.config.Commands, nil
}

func (c *core) RecursivelyCopyCommandDirToPWD(commandName string) error {
	src := fmt.Sprintf("%s/%s/%s.dependencies", EefennCLIRoot, commandName, commandName)
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	script := fmt.Sprintf("%s/%s/%s.sh", EefennCLIRoot, commandName, commandName)
	scriptDst := fmt.Sprintf("%s/%s.sh", pwd, commandName)
	err = copyFile(script, scriptDst)
	if err != nil {
		return err
	}

	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path from the source directory
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(pwd, relPath)

		// If it's a directory, create it
		if d.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// If it's a file, copy it
		return copyFile(path, targetPath)
	})
}

// CopyFile copies a single file from src to dst
func copyFile(src string, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy contents from source to destination
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Ensure file permissions are copied
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

// If the script for the command is in the pwd
func (c *core) RunCommandInPWD(command cmd.CommandInterface) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	scriptPath := fmt.Sprintf("%s/%s", pwd, command)

	// Create the command
	script := exec.Command("/bin/sh", scriptPath)

	script.Stdout = os.Stdout
	script.Stderr = os.Stderr

	err = script.Run()
	if err != nil {
		return err
	}

	return nil
}

// Commit
//
// Add/'commit' a command to core.
func Commit(command commandImageInterface) error {
	var edt cmdFilesController

	// Add the command to the config file
	err := c.config.addCMD(command)
	if err != nil {
		return err
	}
	// Create the directory tree for the command
	err = createCMDImageDir(command)
	if err != nil {
		return err
	}
	// Copy the script for the command from the pwd to the script
	// in newly created directory tree.
	err = c.cmdController.CopyCommand(command)
	if err != nil {
		return err
	}

	err = c.cmdController.CopyDependenciesToDependenciesDir(command)
	if err != nil {
		return err
	}

	err = edt.createCMDImageDir(command)
	if err != nil {
		return err
	}

	return nil
}

func (c *core) RemoveCommandByName(commandName string) error {
	if c == nil {
		return fmt.Errorf("core is not properly initialized\n")
	}

	pCMD, err := c.GetCommandByName(commandName)
	if err != nil {
		return err
	}
	if pCMD == nil {
		return fmt.Errorf("command '%s' does not exist", commandName)
	}

	currentConfig, err := getCurrentConfig()
	if err != nil {
		return err
	}

	err = currentConfig.removeCommandByName(commandName)
	if err != nil {
		return err
	}

	var edt cmdFilesController

	err = edt.RemoveImageDirectoryRecursively(commandName)
	if err != nil {
		return err
	}

	return nil
}
