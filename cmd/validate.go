/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/emam-h/yon/internal/processor"
)

var validateInput string

var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate JSON or YAML file",
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {

		file := validateInput

		// fallback positional
		if file == "" && len(args) == 1 {
			file = args[0]
		}

		if file == "" {
			return fmt.Errorf("no input file provided")
		}

		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if err := processor.Validate(file, data); err != nil {
			return err
		}

		fmt.Println("Valid File")
		return nil
	},
}

func init() {
	validateCmd.Flags().StringVarP(&validateInput, "input", "i", "", "input file")
}