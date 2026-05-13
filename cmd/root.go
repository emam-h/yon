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
	Short: "Convert between YAML and JSON files",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

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
