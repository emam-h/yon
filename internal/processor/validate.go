package processor

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"go.yaml.in/yaml/v4"
)

// Validate checks whether the provided file content is valid JSON or YAML.
// It validates:
//
//   - File extension support
//   - Syntax correctness
//   - Root document structure
//
// Supported extensions:
//
//   - .json
//   - .yaml
//   - .yml
func Validate(file string, data []byte) error {
	ext := normalizeExt(filepath.Ext(file))

	if len(data) == 0 {
		return fmt.Errorf("%s: file is empty", file)
	}
	var v any
	switch ext {
	case ".json":
		if err := validateJSON(data, &v); err != nil {
			return formatJSONError(file, err)
		}

	case ".yaml", ".yml":
		if err := validateYAML(data, &v); err != nil {
			return formatYAMLError(file, err)
		}

	default:
		return fmt.Errorf(
			"%s: unsupported file type %q (supported: .json, .yaml, .yml)",
			file,
			ext,
		)
	}

	if err := validateRootStructure(v); err != nil {
		return fmt.Errorf("%s: %w", file, err)
	}

	return nil
}

func normalizeExt(ext string) string {
	return strings.ToLower(ext)
}

func validateJSON(data []byte, v *any) error {
	dec := json.NewDecoder(strings.NewReader(string(data)))

	// Prevent multiple JSON objects in a single file.
	// Example invalid:
	//
	//   {"a":1} {"b":2}
	//
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		return err
	}

	// Ensure there is no trailing garbage.
	if dec.More() {
		return fmt.Errorf("multiple JSON values detected")
	}

	return nil
}

func validateYAML(data []byte, v *any) error {
	dec := yaml.NewDecoder(strings.NewReader(string(data)))

	if err := dec.Decode(v); err != nil {
		return err
	}

	return nil
}

func validateRootStructure(v any) error {
	switch v.(type) {
	case map[string]any, []any:
		return nil
	default:
		return fmt.Errorf(
			"root document must be an object/map or array/list",
		)
	}
}

func formatJSONError(file string, err error) error {
	switch e := err.(type) {

	case *json.SyntaxError:
		return fmt.Errorf(
			"%s: invalid JSON syntax at byte offset %d",
			file,
			e.Offset,
		)

	case *json.UnmarshalTypeError:
		return fmt.Errorf(
			"%s: invalid JSON type for field %q: expected %v but got %q",
			file,
			e.Field,
			e.Type,
			e.Value,
		)

	default:
		return fmt.Errorf("%s: invalid JSON: %v", file, err)
	}
}

func formatYAMLError(file string, err error) error {
	return fmt.Errorf("%s: invalid YAML: %v", file, err)
}