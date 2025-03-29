package config

import (
	"encoding/json"
	"fmt"
	"github.com/eefenn/eefenn-cli/cmd/subcommand"
	"os"
)

const EefennCLIConfig = "/usr/lib/eefenn-cli/eefenn-cli.config.json"

type subcommandData struct {
	Hash        string `json:"command-hash"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type ConfigObject struct {
	Test subcommandData `json:"test"`
}

// AddCommand
//
// Update /usr/lib/eefenn-cli/eefenn-cli.config.json with
// marshalled subcommand data.
func AddCommand(sc *subcommand.Subcommand) error {
	// get the contents of eefenn-cli.config.json as a map
	configMap, err := getConfigMap()
	if err != nil {
		return err
	}

	// update the map structure, adding the subcommand
	addSubcommandToConfigMap(&configMap, sc)

	// get the data of the updated config map as type []byte
	configByteSlice, err := getConfigByteSliceFromConfigMap(&configMap)
	if err != nil {
		return err
	}

	// write the updated slice of bytes to eefenn-cli.config.json
	err = writeToConfigFile(configByteSlice)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCommand(commandName string) error {
	configMap, err := getConfigMap()
	if err != nil {
		return err
	}

	delete(configMap, commandName)

	configByteSlice, err := getConfigByteSliceFromConfigMap(&configMap)
	if err != nil {
		return err
	}

	// write the updated slice of bytes to eefenn-cli.config.json
	err = writeToConfigFile(configByteSlice)
	if err != nil {
		return err
	}

	return nil
}

// writeToConfigFile
//
// write the contents of a byte array to the eefenn-cli.config.json
func writeToConfigFile(updatedConfig []byte) error {
	// Write directly to file without re-marshaling
	err := os.WriteFile(EefennCLIConfig, updatedConfig, 0666)
	if err != nil {
		return err
	}

	return nil
}

// addSubcommandToConfigMap
//
// update a map of type map[string]interface{} with a subcommand
func addSubcommandToConfigMap(pconfigMap *map[string]interface{}, sc *subcommand.Subcommand) {
	// dereference pointer to unmarshalled JSON
	configMap := *pconfigMap

	// create update the map to include this subcommand as map
	configMap[sc.Name] = map[string]interface{}{
		"description":  sc.Description,
		"command-hash": sc.Hash.String(),
		"script":       fmt.Sprintf("%s.sh", sc.Hash.String()),
	}
}

// getConfigByteSliceFromConfigMap
//
// given a pointer to a map of type, map[string]interface{}, marshal to a
// byte slice
func getConfigByteSliceFromConfigMap(pconfigMap *map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(*pconfigMap, "", "    ")
}

// getConfigMap
//
// Unmarshal the eefenn-cli.config.json file into
// a string:interface map
func getConfigMap() (map[string]interface{}, error) {
	// read the json into byte array
	fileByteArray, err := os.ReadFile(EefennCLIConfig)
	if err != nil {
		return nil, err
	}

	// Step 2: Unmarshal .json into a map
	var unmarshalledJSON map[string]interface{}
	err = json.Unmarshal(fileByteArray, &unmarshalledJSON)
	if err != nil {
		return nil, err
	}

	return unmarshalledJSON, nil
}
