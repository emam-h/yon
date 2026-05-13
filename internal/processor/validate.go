package processor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"go.yaml.in/yaml/v4"
)

func Validate(file string, data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("%s: file is empty", file)
	}

	ext := strings.ToLower(filepath.Ext(file))
	var content any

	switch ext {
	case ".json":
		if err := parseJSON(data, &content); err != nil {
			return formatJSONError(file, data, err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("%s: invalid YAML: %w", file, err)
		}
	default:
		return fmt.Errorf("%s: unsupported file extension %q", file, ext)
	}

	if err := validateRoot(content); err != nil {
		return fmt.Errorf("%s: %w", file, err)
	}

	return nil
}

func parseJSON(data []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(target); err != nil {
		return err
	}
	if decoder.More() {
		return fmt.Errorf("extra data after top-level JSON value")
	}
	return nil
}

func validateRoot(value any) error {
	switch value.(type) {
	case map[string]any, []any:
		return nil
	default:
		return fmt.Errorf("root must be an object or array, got %T", value)
	}
}

func formatJSONError(file string, data []byte, err error) error {
	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		line, col := lineCol(data, syntaxErr.Offset)
		lineText := getLine(data, line)
		caret := " " + strings.Repeat(" ", col-1) + "^"
		return fmt.Errorf(
			"%s: invalid JSON syntax at line %d, column %d\n%s\n%s\n%w",
			file,
			line,
			col,
			lineText,
			caret,
			syntaxErr,
		)
	}
	return fmt.Errorf("%s: invalid JSON: %w", file, err)
}

func lineCol(data []byte, offset int64) (line, col int) {
	if offset < 1 {
		offset = 1
	}

	line = 1
	col = 1

	for i := int64(0); i < offset-1 && i < int64(len(data)); i++ {
		if data[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return
}

func getLine(data []byte, target int) string {
	line := 1
	start := 0

	for i, b := range data {
		if b == '\n' {
			if line == target {
				return string(data[start:i])
			}
			line++
			start = i + 1
		}
	}

	if line == target {
		return string(data[start:])
	}
	return ""
}
