package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/woven-planet/go-zserio/internal/generator"
	"github.com/woven-planet/go-zserio/internal/model"
)

var (
	outputDirectory string
	topLevelPackage string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates Go files for the source zserio files",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := model.FromFilesystem(args[0])
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Parse error: %v", err)
			return err
		}
		fmt.Printf("Parsing complete, generating files...\n")
		if err = generator.Generate(m, strings.TrimSpace(outputDirectory), &generator.Options{
			RootPackage: strings.TrimSpace(topLevelPackage),
		}); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Code generation error: %v", err)
			return err
		}
		fmt.Printf("Code generation completed.\n")
		return nil
	},
}

func init() {
	generateCmd.Flags().StringVarP(&outputDirectory, "out", "o", "", "Output directory to generate Go files to")
	generateCmd.MarkFlagRequired("out")

	generateCmd.Flags().StringVarP(&topLevelPackage, "rootpackage", "r", "", "Name of the root package package")
	generateCmd.MarkFlagRequired("rootpackage")

	rootCmd.AddCommand(generateCmd)
}
