// asc.go
//
// asc (add subcommand)

package asc

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
)

const EefennCLIDir = "/usr/lib/eefenn-cli"

type Subcommand struct {
	Name  string          `json:"name"`
	Entry SubcommandEntry `json:"entry"`
}
type SubcommandEntry struct {
	// unique identifier for the command
	Hash uuid.UUID `json:"command-hash"`

	// the script which the command is an alias for
	SourceScript string `json:"script"`

	// script dependencies
	Dependencies []string `json:"dependencies,omitempty"`

	// description for what the script does
	Description string `json:"description,omitempty"`
}

// CreateSubCommand
//
// Create a subcommand struct based on required command information
func CreateSubCommand(name string, sourceScriptName string, dependencyPaths []string, description string) Subcommand {
	UUID := uuid.New()
	subCommandEntry := SubcommandEntry{
		Hash:         UUID,
		SourceScript: sourceScriptName,
		Dependencies: dependencyPaths,
		Description:  description,
	}
	subCommand := Subcommand{
		Name:  name,
		Entry: subCommandEntry,
	}
	return subCommand
}

func (sc *Subcommand) CreateSubCommandConfigEntry() error {
	commandJson, err := sc.GetSubCommandJson()
	if err != nil {
		return err
	}

	fmt.Printf(string(commandJson))

	return nil
}

func (sc *Subcommand) GetSubCommandJson() ([]byte, error) {
	jsonData, err := json.MarshalIndent(sc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON:", err)
	}

	return jsonData, nil
}

// CreateSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the subcommand
func (sc *Subcommand) CreateSubcommandDirTree() error {
	// create the directory that contains dependencies and script for the command
	subCommandDir := sc.GetSubcommandDependenciesDirectory()

	err := os.MkdirAll(subCommandDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this subcommand: %v\n", err)
	}

	// create a blank command script
	blankFile, err := sc.CreateEmptySubcommandShellFile()
	if err != nil {
		return fmt.Errorf("Could not create empty subcommand .sh file\n")
	}

	// write the contents of the command script to the persisted script
	_, err = blankFile.Write([]byte("hello"))
	if err != nil {
		return fmt.Errorf("Failed to copy the contennts of the target shell script\n")
	}

	return nil
}

// GetAbsoluteSubcommandDirname
//
// get the absolute directory path for the subcommand directory.
func (sc *Subcommand) GetAbsoluteSubcommandDirname() string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", EefennCLIDir, sc.Entry.Hash.String())

	return commandDirectory
}

// GetSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func (sc *Subcommand) GetSubcommandDependenciesDirectory() string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s.dependencies", EefennCLIDir, sc.Entry.Hash.String())

	return commandDependenciesDirectory
}

func (sc *Subcommand) CreateEmptySubcommandShellFile() (*os.File, error) {
	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s%s", sc.Entry.Hash.String(), ".sh")

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
