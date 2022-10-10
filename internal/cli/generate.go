package cli

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/woven-planet/go-zserio/internal/generator"
	"github.com/woven-planet/go-zserio/internal/model"
)

var (
	outputDirectory string
	topLevelPackage string
	noFormat        bool
	emitSQLSupport  bool
	maxPathLength   int
	outputPackage   string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates Go files for the source zserio files",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := model.FromFilesystem(args...)
		if err != nil {
			return fmt.Errorf("parse schema: %w", err)
		}
		log.Println("Parsing complete, generating files...")
		var options []generator.Option
		if noFormat {
			options = append(options, generator.DoNotFormatSource)
		}
		if emitSQLSupport {
			options = append(options, generator.EmitSQLSupport)
		}
		if maxPathLength > 0 {
			options = append(options, generator.PathLengthLimiter{MaxPathLength: maxPathLength}.LimitPathLength)
		}

		if err = generator.Generate(m,
			strings.TrimSpace(outputDirectory),
			strings.TrimSpace(topLevelPackage),
			strings.TrimSpace(outputPackage),
			options...,
		); err != nil {
			return fmt.Errorf("generate code: %w", err)
		}
		log.Println("Code generation completed.")
		return nil
	},
}

func init() {
	generateCmd.Flags().StringVarP(&outputDirectory, "out", "o", "", "Output directory to generate Go files to. If it is '-' then we will print a single package to stdout.")
	generateCmd.MarkFlagRequired("out")

	generateCmd.Flags().StringVarP(&topLevelPackage, "rootpackage", "r", "", "Name of the root package package")
	generateCmd.MarkFlagRequired("rootpackage")

	generateCmd.Flags().StringVarP(&outputPackage, "only", "", "", "Output only a single package.")

	generateCmd.Flags().BoolVarP(&noFormat, "noformat", "", false, "Do not format the source files.")
	generateCmd.Flags().MarkHidden("noformat")

	generateCmd.Flags().BoolVar(&emitSQLSupport, "sql", false, "Emit code for SQL support")

	generateCmd.Flags().IntVar(&maxPathLength, "max-path-length", generator.DefaultMaxPathLength, "Maximum length of generated file paths. Set to 0 for no limit.")

	rootCmd.AddCommand(generateCmd)
}
