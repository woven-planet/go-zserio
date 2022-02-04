package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/woven-planet/go-zserio/internal/generator"
	"github.com/woven-planet/go-zserio/internal/model"
)

var testCmd = &cobra.Command{
	Use:  "test",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := model.FromFilesystem(args[0])
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Parse error: %v", err)
			return err
		}

		if err = generator.Generate(m, "/tmp/code", &generator.Options{
			RootPackage:       "testpackage",
			DoNotFormatSource: false,
		}); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Code generation error: %v", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

/*

// UnmarshalZserio implements the zserio.Unmarshaler interface.
func (v *{{ $.enum.Name }}) UnmarshalZserio(r *bitio.Reader) error {
    var err error
    {{- if not $enum.Type.IsBuiltin }}
      panic("only builtin types supported")
    {{- else if eq $enum.Type.Bits 0 }}
      *v, err = ztype.Read{{ title $enum.Type.Name }}(w, v)
    {{- else if eq $enum.Type.Name "bit" }}
      v, err = ztype.ReadUnsignedBitfield(w, v, {{ $enum.Type.Bits }})
    {{- else if eq $enum.Type.Name "int" }}
      v, err = ztype.ReadSignedBitfield(w, v, {{ $enum.Type.Bits }})
    {{- else }}
      panic("unsupported type")
    {{- end }}
    return err
}

*/
