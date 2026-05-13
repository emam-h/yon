# yon

A fast and simple CLI tool for converting and processing YAML and JSON files. Convert between formats, validate syntax, and pretty-print your data files with ease.

## Features

- **Convert**: Seamlessly convert between YAML and JSON formats
- **Validate**: Validate JSON and YAML files with detailed error reporting including line and column numbers
- **Format**: Pretty-print and reformat JSON and YAML files for consistent output
- **Cross-platform**: Available for macOS, Linux, and Windows (both ARM64 and AMD64 architectures)

## Installation

### Homebrew (macOS/Linux)

```bash
brew tap emam-h/tools
brew install yon
```

### From Source

```bash
git clone https://github.com/emam-h/yon.git
cd yon
go install
```

### Build Locally

```bash
go build -o yon
```

Or build for specific platforms:

```bash
make build
```

This creates binaries in the `dist/` directory for:
- macOS (ARM64 and x86_64)
- Linux (ARM64 and x86_64)
- Windows (ARM64 and x86_64)

## Usage

### Convert Between Formats

Convert a YAML file to JSON or vice versa:

```bash
# Using positional arguments
yon input.yaml output.json
yon input.json output.yaml

# Using flags
yon --input=input.yaml --output=output.json
yon -i input.json -o output.yaml
```

### Format Files

Pretty-print and reformat JSON or YAML files:

```bash
# Format in-place (overwrites the original file)
yon format file.json
yon format file.yaml

# Format to a different output file
yon format input.json output.json
yon format input.yaml output.yaml

# Using flags
yon format --input=file.json --output=formatted.json
yon format -i file.yaml -o formatted.yaml
```

### Validate Files

Validate JSON or YAML syntax with detailed error messages:

```bash
# Validate using positional argument
yon validate file.json
yon validate file.yaml

# Validate using flag
yon validate --input=file.json
yon validate -i file.yaml
```

**Output:**
- ✓ If valid: `Valid File`
- ✗ If invalid: Error message with file path, line number, column number, and error description

Example error output:
```
file.json: invalid JSON syntax at line 2, column 15: invalid character '}' looking for beginning of value
```

## Examples

### Example 1: Convert and Validate

```bash
# Convert YAML to JSON
yon config.yaml config.json

# Validate the converted file
yon validate config.json

# Pretty-print the result
yon format config.json
```

### Example 2: Batch Processing

```bash
# Format all JSON files in a directory
for file in *.json; do
  yon format "$file"
done

# Validate all YAML files
for file in *.yaml; do
  yon validate "$file"
done
```

### Example 3: Pipeline Integration

```bash
# Convert YAML config to JSON, then validate
yon input.yaml output.json && yon validate output.json

# Format a file and then convert it
yon format data.json && yon data.json data.yaml
```

## Supported File Types

- **JSON**: `.json` files
- **YAML**: `.yaml` and `.yml` files

## License

Copyright © 2026 Emam Hasan. Licensed under the terms in the LICENSE file.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests on GitHub.