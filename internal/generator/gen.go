package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

type Options struct {
	RootPackage       string
	DoNotFormatSource bool
	OutputPackage     string

	outputToStdout bool
}

type data map[string]interface{}

func Generate(m *model.Model, rootPath string, options *Options) error {
	options.outputToStdout = rootPath == "-"

	if options.outputToStdout {
		pkg, ok := m.Packages[options.OutputPackage]
		if !ok {
			var known []string
			for name := range m.Packages {
				known = append(known, name)
			}
			sort.Strings(known)

			return fmt.Errorf("%q not found, only know about %q", options.OutputPackage, known)
		}

		return generatePackage(rootPath, pkg, options)
	}

	for _, pkg := range m.Packages {
		if err := generatePackage(rootPath, pkg, options); err != nil {
			return err
		}
	}
	return nil
}

func executeTemplate(out io.Writer, tmpl string, data data) error {
	templ := templates.Lookup(tmpl)
	if templ == nil {
		return fmt.Errorf("template %q is missing", tmpl)
	}

	return templ.Execute(out, data)
}

func writeToFile(code []byte, rootPath, zserioPkg, fn string, doNotFormatCode bool) (retErr error) {
	ids := strings.Split(strings.ToLower(zserioPkg), ".")
	outputDir := path.Join(append([]string{rootPath}, ids...)...)
	outputFile := path.Join(outputDir, fn)

	log.Printf("Writing %s\n", outputFile)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	f, err := os.OpenFile(outputFile, FileOpenFlags, 0644)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer func() {
		if retErr == nil {
			return
		}

		_ = os.Remove(outputFile)
	}()
	defer f.Close()

	return writeGoSource(f, code, doNotFormatCode)
}

func writeGoSource(w io.Writer, code []byte, doNotFormatCode bool) error {
	if doNotFormatCode {
		if _, err := w.Write(code); err != nil {
			return fmt.Errorf("write: %w", err)
		}
		return nil
	}

	code, err := StripImports(code)
	if err != nil {
		return fmt.Errorf("strip imports: %w", err)
	}

	formatted, err := format.Source(code)
	if err != nil {
		log.Println(string(code))
		return fmt.Errorf("format: %w", err)
	}

	if _, err := w.Write(formatted); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	return nil
}

// output allows us easier passing of arguments to write packages.
type output struct {
	name         string // the name of the data structure that is output
	template     string // template file name
	data         data
	withPreamble bool
}

func newOutput(name, template string, data data, withPreamble bool) output {
	return output{
		name:         name,
		template:     template,
		data:         data,
		withPreamble: withPreamble,
	}
}

func newTypeOutput(astType interface{}, d data, withPreamble bool) output {
	var name, typeName string
	switch t := astType.(type) {
	case *ast.Enum:
		name = t.Name
		typeName = "enum"
	case *ast.Union:
		name = t.Name
		typeName = "union"
	case *ast.BitmaskType:
		name = t.Name
		typeName = "bitmask"
	case *ast.Choice:
		name = t.Name
		typeName = "choice"
	case *ast.Struct:
		name = t.Name
		typeName = "struct"
	default:
		panic(fmt.Sprintf("%#T is unsupported", t))
	}

	templateData := make(data, len(d)+1)
	for k, v := range d {
		templateData[k] = v
	}
	templateData[typeName] = astType
	return newOutput(strcase.ToSnake(name)+".go", typeName+".go.tmpl", templateData, withPreamble)
}

func (o *output) executeTemplate(w io.Writer) error {
	if o.withPreamble {
		if err := executeTemplate(w, "preamble.go.tmpl", o.data); err != nil {
			return fmt.Errorf("generate preamble: %w", err)
		}
	}

	return executeTemplate(w, o.template, o.data)
}

func newOutputs(pkg *ast.Package, rootPackage string, singleFile bool) []output {
	data := data{
		"pkg":         pkg,
		"rootPackage": rootPackage,
	}

	outs := []output{newOutput("pkg.go", "package.go.tmpl", data, true)}

	for _, t := range pkg.Enums {
		outs = append(outs, newTypeOutput(t, data, !singleFile))
	}

	for _, t := range pkg.Unions {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !singleFile))
	}

	for _, t := range pkg.Bitmasks {
		outs = append(outs, newTypeOutput(t, data, !singleFile))
	}

	for _, t := range pkg.Choices {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !singleFile))
	}

	for _, t := range pkg.Structs {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !singleFile))
	}

	return outs
}

func generatePackage(rootPath string, pkg *ast.Package, options *Options) error {
	outs := newOutputs(pkg, options.RootPackage, options.outputToStdout)

	if options.outputToStdout {
		var out bytes.Buffer
		for _, o := range outs {
			if err := o.executeTemplate(&out); err != nil {
				return fmt.Errorf("generate %s: %w", o.name, err)
			}
		}

		return writeGoSource(os.Stdout, out.Bytes(), options.DoNotFormatSource)
	}

	for _, o := range outs {
		var out bytes.Buffer
		if err := o.executeTemplate(&out); err != nil {
			return fmt.Errorf("generate %s: %w", o.name, err)
		}

		if err := writeToFile(out.Bytes(), rootPath, pkg.Name, o.name, options.DoNotFormatSource); err != nil {
			return fmt.Errorf("write %s: %w", o.name, err)
		}
	}

	return nil
}
