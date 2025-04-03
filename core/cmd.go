package core

import (
	"fmt"
	"os"
	"os/exec"
)

const EefennCLIRoot = "/usr/lib/eefenn-cli"

type absolutePath string

type commandInstanceInterface interface {
	CreateCMDInstanceAtPosition(position absolutePath) RunnableCommandInterface

	// CleanUp
	// Delete all files of the CommandImage
	CleanupInstance() error

	// ChangeRootPosition
	// Change the position of the CommandImage image in the filesystem
	MoveCommand(newRootPosition absolutePath) CommandInterface
}

type commandImageInterface interface {
	CreateRunnableCommand(commandLocation absolutePath) RunnableCommandInterface

	NewImageDirectory() (CommandInterface, error)

	RemoveImageDirectory() error

	// Gets the directory path for a command image
	GetImageDirPath() string

	GetImageFileContents(fileName string) ([]byte, error)

	// GetName
	// Get the name of the CommandImage
	GetName() string

	// GetCommandFilePathsRelativeToCMDRoot
	// Get the paths of the CommandImage relative to the root of the CommandImage
	GetCommandFilePathsRelativeToCMDRoot() ([]relativeFilePath, error)
}

type relativeFilePath string

type RunnableCommandInterface interface {
	Run() error
}

type CommandImage struct {
	// alias of the script
	name string `yaml:"name" json:"name"`

	// the script which the CommandImage is an alias for
	script string `yaml:"script" json:"script"`

	// array of filepaths that the script needs to run
	dependencies []string `yaml:"dependencies" json:"dependencies"`

	// description for what the script does
	description string `yaml:"description" json:"description"`

	// the arguments to the CommandImage
	args []Arg ` yaml:"args" json:"args"`

	// Position of the root of the CommandImage (where the script is)
	// relative to the root of the filesystem
	rootPosition string
}

func (cmd *CommandImage) getImageDirPath() string {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) getImageDirectoryPosition() string {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) GetCommandFilePathsRelativeToCMDRoot() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) CleanUp() error {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) ChangeRootPosition(newPosition string) CommandImage {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) CreateImageDirectory() error {
	dirname := cmd.getImageDirPath()

	err := os.RemoveAll(dirname)
	if err != nil {
		return err
	}
}

func (cmd *CommandImage) RemoveImageDirectory() error {
	//TODO implement me
	panic("implement me")
}

func (cmd *CommandImage) Write() error {
	files, err := cmd.GetCommandFilePathsRelativeToCMDPosition()
	if err != nil {
		return err
	}

	for index, file := range files {
		// create the file
		newFile, err := os.Create(file)
		if err != nil {
			return err
		}

		contentsOfImageFile, err := os.ReadFile(cmd.getImageFileContents(file))
		if err != nil {
			return err
		}

		newFile.Write(contentsOfImageFile)
	}

	return nil
}

func (cmd *CommandImage) getImageFileContents() []byte {
	imageDirPath := cmd.getImageDirPath()
	imageFilePath := fmt.Sprintf("%s"/"%s", imageDirPath)

	os.ReadFile(imageDirPath)
}

func (cmd *CommandImage) Run() error {
	scriptCMD := exec.Command("sh", cmd.script)

	return scriptCMD.Run()
}

type Arg struct {
	Name        string `json:"name" json:"name"`
	Type        string `json:"type" json:"type"`
	Description string `json:"description" json:"description"`
}

// GetCommandFilePaths
// Gets the file paths for script, configuration file and all dependencies
// and returns these paths relative to where the CommandImage currently is.
func (cmd *CommandImage) GetCommandFilePathsRelativeToCMDPosition() ([]string, error) {
	var commandFilePaths []string

	// add CommandImage script
	commandFilePaths = append(commandFilePaths, cmd.script)

	// add config.yaml

	// add all dependency files
	for _, dependency := range cmd.dependencies {
		commandFilePaths = append(commandFilePaths, dependency)
	}

	return nil, nil
}

// GetCommandFilePaths
// Gets the file paths for script, configuration file and all dependencies
// and returns these paths relative to where the CommandImage currently is.
func (cmd *CommandImage) GetName() string {
	return cmd.name
}

func getImageDirPath(cmd CommandInterface) string {
	return fmt.Sprintf("%s/%s", EefennCLIRoot, cmd.GetName())
}
