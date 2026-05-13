package processor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"go.yaml.in/yaml/v4"
)

func Format(input, output string) error {

	data, err := os.ReadFile(input)
	if err != nil {
		return fmt.Errorf("Failed to read file: %w", err)
	}

	formatted, err := formatData(input, data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(output, formatted, 0644); err != nil {
		return fmt.Errorf("Failed to write file: %w", err)
	}

	return nil
}

func formatData(filename string, data []byte) ([]byte, error) {

	ext := filepath.Ext(filename)
	var v any
	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("Invalid JSON: %w", err)
		}
		res, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("Failed to format JSON: %w", err)
		}
		return res, nil

	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &v); err != nil {
			return nil, fmt.Errorf("Invalid YAML: %w", err)
		}
		res, err := yaml.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("Failed to format YAML: %w", err)
		}
		return res, nil

	default:
		return nil, fmt.Errorf(
			"Unsupported file extension: %s",
			ext,
		)
	}
}