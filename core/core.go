// core.go
//
// Core is essentially the management system for commands, and their files.
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
	"strings"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

// GetAllCommands
// Gets all commands in config. If there are no commands,
// will return an error
func GetAllCommands() ([]cmd.Command, error) {
	config, err := getCurrentConfig()
	if err != nil {
		return nil, err
	}

	if len(config.Commands) == 0 {
		return nil, fmt.Errorf("there are no commands")
	}
	return config.Commands, nil
}

// Commit
//
// Add/'commit' a command to core.
func Commit(command cmd.Command) error {
	_, pCMD, err := GetCommandByName(command.Name)
	if pCMD != nil {
		return fmt.Errorf("command '%s' already exists\n\nUse the 'ef rm' command to remove this command, or 'ef edit' to edit the command.", command.Name)
	}

	config, err := getCurrentConfig()
	if err != nil {
		return err
	}

	// Add the command to the config file
	err = config.addCMD(command)
	if err != nil {
		return err
	}
	// Create the directory tree for the command

	return nil
}

func CreateCommandInDir(commandName string, dirPath string) error {
	// get the map of absolute paths from image directory
	// file to the run directory file (<image_dir/file_x>:<run_dir/file_x>)
	imgFilesRunFilesMap, err := getImgFilesToRunFilesMap(commandName, dirPath)
	if err != nil {
		return err
	}

	// initialize a slice of absolute paths to keep track
	// of the files we have copied from image directory for
	// cleaning up
	var runFiles []string
	// copy all files in the image to the run directory
	for source, destination := range imgFilesRunFilesMap {
		err := copyFile(source, destination)
		runFiles = append(runFiles, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func RemoveCommandByName(commandName string) error {
	_, pCMD, err := GetCommandByName(commandName)
	if err != nil {
		return err
	}
	if pCMD == nil {
		return fmt.Errorf("command '%s' does not exist", commandName)
	}

	// get an object to represent the current config file
	currentConfig, err := getCurrentConfig()
	if err != nil {
		return err
	}

	// remove the command from the config file
	err = currentConfig.removeCommandByName(commandName)
	if err != nil {
		return err
	}

	// remove the directory storing the image for the command
	err = RemoveCommandImgDir(commandName)
	if err != nil {
		return err
	}

	return nil
}

func Run(cmdName string, runDir string) error {
	// get the map of absolute paths from image directory
	// file to the run directory file (<image_dir/file_x>:<run_dir/file_x>)
	imgFilesRunFilesMap, err := getImgFilesToRunFilesMap(cmdName, runDir)
	if err != nil {
		return err
	}

	// initialize a slice of absolute paths to keep track
	// of the files we have copied from image directory for
	// cleaning up
	var runFiles []string
	// copy all files in the image to the run directory
	for source, destination := range imgFilesRunFilesMap {
		err := copyFile(source, destination)
		runFiles = append(runFiles, destination)
		if err != nil {
			return err
		}
	}

	_, command, err := GetCommandByName(cmdName)
	if err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	// get the absolute path to the command script in the run directory
	runScript := fmt.Sprintf("%s/%s", pwd, command.Script)

	// run the script
	script := exec.Command(runScript)
	err = script.Run()
	if err != nil {
		return err
	}

	return nil
}

func cleanupFiles(filesToRemove []string) {
	for _, file := range filesToRemove {
		err := os.Remove(file)
		if err != nil {
			fmt.Printf("unable to remove file: %s", file)
			continue
		}
	}

	return
}

// CopyFile copies a single file from src to dst
func copyFile(src string, dst string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Printf("could not close file %s", srcFile.Name())
			return
		}
	}(srcFile)

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Printf("could not close file %s", srcFile.Name())
		}
	}(dstFile)

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

func getImgFilesToRunFilesMap(cmdName string, runDir string) (map[string]string, error) {
	runDirPathCleaned := fmt.Sprintf("%s%s", strings.TrimRight(runDir, "/"), "/")

	paths, err := getCommandImgAbsFilePaths(cmdName)

	if err != nil {
		return nil, err
	}

	imgFilesRunFilesMap := make(map[string]string)
	cmdImgDirPath := EefennCLIRoot + "/" + cmdName

	for _, file := range paths {
		imageRelPath := strings.Replace(file, cmdImgDirPath, "", 1)
		fmt.Printf("%s\n", runDirPathCleaned)
		imgFilesRunFilesMap[file] = prependPath(runDirPathCleaned, strings.TrimPrefix(imageRelPath, "/"))
	}

	return imgFilesRunFilesMap, nil
}

func prependPath(pathToPrepend string, pathToPrependTo string) string {
	fmt.Printf("pathToPrepend: %s\n", pathToPrepend)
	fmt.Printf("pathToPrependTo: %s\n", pathToPrependTo)
	return pathToPrepend + "/" + pathToPrependTo
}

func getCommandImgAbsFilePaths(commandName string) ([]string, error) {
	var filepaths []string

	err := filepath.WalkDir(EefennCLIRoot+"/"+commandName, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filepaths = append(filepaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filepaths, nil
}

// RemoveCommandImgDir
//
// remove a command directory recursively by command hash
func RemoveCommandImgDir(commandName string) error {
	_, command, err := GetCommandByName(commandName)
	if err != nil {
		return err
	}

	err = os.RemoveAll(command.GetCmdImgDirPath())
	if err != nil {
		return err
	}

	return nil
}
