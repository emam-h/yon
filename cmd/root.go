/*
Copyright © 2026 Emam Hasan
*/
package cmd

import (
	"errors"
	"github.com/emam-h/yon/internal/processor"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string

var rootCmd = &cobra.Command{
	Use:   "yon",
	Short: "Convert files between YAML and JSON formats",
	Long: `Yon is a CLI tool for converting data files between YAML and JSON formats.
It accepts either positional source and destination file arguments or the
--input and --output flags. Yon also provides subcommands to validate and
reformat YAML/JSON content for consistent output.

Examples:
  yon input.yaml output.json
  yon --input=input.json --output=output.yaml
  yon format file.yaml
  yon validate file.json`,

	RunE: func(cmd *cobra.Command, args []string) error {

		if inputFile != "" && outputFile != "" {

			return processor.Convert(inputFile, outputFile)
		}
		if len(args) == 2 {
			return processor.Convert(args[0], args[1])
		}
		return errors.New(
			"usage:\n" +
				"  yon input.yaml output.json\n" +
				"  yon -i input.yaml -o output.json",
		)
	},
}

func init() {
	rootCmd.Flags().StringVarP(
		&inputFile,
		"input",
		"i",
		"",
		"Input file",
	)

	rootCmd.Flags().StringVarP(
		&outputFile,
		"output",
		"o",
		"",
		"Output file",
	)
	rootCmd.AddCommand(formatCmd)
	rootCmd.AddCommand(validateCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
