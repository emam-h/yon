package processor

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"go.yaml.in/yaml/v4"
)

func Validate(file string, data []byte) error {

	ext := filepath.Ext(file)
	var v any
	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("invalid JSON: %w", err)
		}
		return nil
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &v); err != nil {
			return fmt.Errorf("invalid YAML: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("Unsupported file extension: %s", ext)
	}
}