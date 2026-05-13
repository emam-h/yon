package processor

import (
	"fmt"
	"os"
	"errors"
	"path/filepath"
	"encoding/json"
	"go.yaml.in/yaml/v4"
)

func Convert(inputFile string, outputFile string) error {

	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	inputExt := filepath.Ext(inputFile)
	outputExt := filepath.Ext(outputFile)

// Yaml to JSON
	if (inputExt == ".yaml" || inputExt == ".yml") &&
		outputExt == ".json" {

		var obj interface{}

		err = yaml.Unmarshal(data, &obj)
		if err != nil {
			return err
		}

		jsonData, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return err
		}

		err = os.WriteFile(outputFile, jsonData, 0644)
		if err != nil {
			return err
		}

		fmt.Printf("Converted %s -> %s\n", inputFile, outputFile)
		return nil
	}
// JSON to Yaml
	if inputExt == ".json" &&
		(outputExt == ".yaml" || outputExt == ".yml") {

		var obj interface{}

		err = json.Unmarshal(data, &obj)
		if err != nil {
			return err
		}

		yamlData, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}

		err = os.WriteFile(outputFile, yamlData, 0644)
		if err != nil {
			return err
		}

		fmt.Printf("Converted %s -> %s\n", inputFile, outputFile)
		return nil
	}

	return errors.New("Unsupported conversion")
}