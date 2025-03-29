// add_subcommand.go
//
// asc (add subcommand) is a method of customizing the command line tool
// by updating the directory /usr/lib/eefenn-cli and eefenn-cli.config.json
//
// @author Mikey Fennelly

package add_subcommand

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"os"
)

const eefennCLIDir = "/usr/lib/eefenn-cli"
const configJSONPath = eefennCLIDir + "/eefenn-cli.config.json"

var AscCommand = &cobra.Command{
	Use:   "asc",
	Short: "Add a subcommand to eefenn-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding subcommand...")
	},
}

type subcommand struct {
	// alias of the script
	Name string `json:"name"`

	// unique identifier for the command
	Hash uuid.UUID `json:"command-hash"`

	// the script which the command is an alias for
	SourceScript string `json:"script"`

	// script dependencies
	Dependencies []string `json:"dependencies,omitempty"`

	// description for what the script does
	Description string `json:"description,omitempty"`
}

// AddSubCommand
//
// Add a subcommand, and it's script to the user's CLI
func AddSubCommand(name string, sourceScriptName string, dependencyPaths []string, description string) error {
	subCommand := createSubCommand(name, sourceScriptName, dependencyPaths, description)

	err := subCommand.createSubCommandConfigEntry()
	if err != nil {
		return err
	}

	err = subCommand.createSubcommandDirTree()
	if err != nil {
		return err
	}

	return nil
}

// createSubCommand
//
// Create a subcommand struct based on required command information
func createSubCommand(name string, sourceScriptName string, dependencyPaths []string, description string) subcommand {
	UUID := uuid.New()
	subCommand := subcommand{
		Name:         name,
		Hash:         UUID,
		SourceScript: sourceScriptName,
		Dependencies: dependencyPaths,
		Description:  description,
	}
	return subCommand
}

// createSubCommandConfigEntry
//
// Create a JSON object for this subcommand and update
// eefenn-cli.config.json with that object.
func (sc *subcommand) createSubCommandConfigEntry() error {
	commandJson, err := sc.getSubCommandJson()
	if err != nil {
		return err
	}

	err = os.WriteFile(configJSONPath, commandJson, 0666)
	if err != nil {
		return err
	}

	fmt.Printf(string(commandJson))

	return nil
}

// getSubCommandJson
//
// Get the JSON object for this command.
func (sc *subcommand) getSubCommandJson() ([]byte, error) {
	jsonData, err := json.MarshalIndent(sc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %v\n", err)
	}

	return jsonData, nil
}

// createSubcommandDirTree
//
// Create an entry in /usr/lib/eefenn-cli for the subcommand
func (sc *subcommand) createSubcommandDirTree() error {
	// create the directory that contains dependencies and script for the command
	subCommandDependenciesDir := sc.getSubcommandDependenciesDirectory()

	err := os.MkdirAll(subCommandDependenciesDir, 0755)
	if err != nil {
		return fmt.Errorf("Could not create directory for this subcommand: %v\n", err)
	}

	subCommandDir := sc.getAbsoluteSubcommandDirname()
	// create a blank command script
	blankFile, err := sc.createEmptySubcommandShellFile(subCommandDir)
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

// getAbsoluteSubcommandDirname
//
// get the absolute directory path for the subcommand directory.
func (sc *subcommand) getAbsoluteSubcommandDirname() string {
	// create the string for the command ID
	commandDirectory := fmt.Sprintf("%s/%s", eefennCLIDir, sc.Hash.String())

	return commandDirectory
}

// getSubcommandDependenciesDirectory
//
// Get the file path to /usr/lib/eefenn-cli/<command-hash>/<command-hash>.dependencies
func (sc *subcommand) getSubcommandDependenciesDirectory() string {
	// create the string for the command ID
	commandDependenciesDirectory := fmt.Sprintf("%s/%s/%s.dependencies", eefennCLIDir, sc.Hash.String(), sc.Hash.String())

	return commandDependenciesDirectory
}

// createEmptySubcommandShellFile
//
// Create an empty shell file of the name <command-hash>.sh
func (sc *subcommand) createEmptySubcommandShellFile(parentDir string) (*os.File, error) {
	// create '<command-hash>.sh' filename string
	fileName := fmt.Sprintf("%s/%s%s", parentDir, sc.Hash.String(), ".sh")

	// create the file
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// updateConfigJSON
//
// Update /usr/lib/eefenn-cli/effen-cli.config.json with
// marshalled subcommand data.
func (sc *subcommand) updateConfigJSON() error {
	configJSON, err := os.ReadFile(configJSONPath)
	if err != nil {
		return err
	}

	// Step 2: Unmarshal eefenn-cli.config.json into a map
	var data map[string]interface{}
	err = json.Unmarshal(configJSON, &data)
	if err != nil {
		return err
	}

	data[sc.Name] = map[string]interface{}{
		"description":  sc.Description,
		"command-hash": sc.Hash.String(),
		"script":       fmt.Sprintf("%s.sh", sc.Hash.String()),
		"dependencies": sc.Dependencies,
	}

	updatedConfig, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(configJSONPath, updatedConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}
