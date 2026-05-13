/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/emam-h/yon/internal/processor"
)

var formatInput string
var formatOutput string

var formatCmd = &cobra.Command{
	Use:   "format [input] [output]",
	Short: "Pretty print JSON or YAML",

	RunE: func(cmd *cobra.Command, args []string) error {

		in := formatInput
		out := formatOutput

		if in != "" {

			if out == "" {
				out = in
			}

			return processor.Format(in, out)
		}

		switch len(args) {
			case 1:

				in = args[0]
				out = args[0]

			case 2:

				in = args[0]
				out = args[1]

			default:
				return fmt.Errorf(
					"usage:\n" +
						"  yon format file.yaml\n" +
						"  yon format input.yaml output.yaml\n" +
						"  yon format -i input.yaml -o output.yaml",
				)
		}

		return processor.Format(in, out)
	},
}

func init() {
	formatCmd.Flags().StringVarP(&formatInput, "input", "i", "", "input file")
	formatCmd.Flags().StringVarP(&formatOutput, "output", "o", "", "output file")
}