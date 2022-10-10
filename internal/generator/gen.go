package generator

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"

	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

type options struct {
	rootPackage       string
	doNotFormatSource bool
	outputPackage     string
	outputToStdout    bool
	EmitSQLSupport    bool
	maxPathLength     int
}

const (
	DefaultMaxPathLength = 260 // Hard limit on Windows
	MinFileNameLength    = 2
	FileSuffixLength     = 5 // Leave space for at least single-digit index if needed
	FileSuffix           = ".go"
)

// Option sets a configuration option
type Option func(opts *options)

// DoNotFormatSource indicates generate Go source should not be formatted
func DoNotFormatSource(opts *options) {
	opts.doNotFormatSource = true
}

// EmitSQLSupport tells the generator to produce support for storing enum
// types in a SQL database.
func EmitSQLSupport(opt *options) {
	opt.EmitSQLSupport = true
}

// PathLengthLimiter indicates that generated source file paths should be clipped to the given length
type PathLengthLimiter struct {
	MaxPathLength int
}

func (limiter PathLengthLimiter) LimitPathLength(opts *options) {
	opts.maxPathLength = limiter.MaxPathLength
}

type data map[string]any

// uniquePaths keeps track of all codes files that have been generated so far and adds an "_idx" suffix if the same path appears again
var uniquePaths map[string]int

func Generate(m *model.Model, rootPath, rootPackage, outputPackage string, flags ...Option) error {
	opts := &options{
		outputToStdout: rootPath == "-",
		rootPackage:    rootPackage,
	}

	for _, f := range flags {
		f(opts)
	}

	if opts.outputToStdout {
		pkg, ok := m.Packages[outputPackage]
		if !ok {
			var known []string
			for name := range m.Packages {
				known = append(known, name)
			}
			sort.Strings(known)

			return fmt.Errorf("%q not found, only know about %q", outputPackage, known)
		}

		return generatePackage(rootPath, pkg, opts)
	}

	for _, pkg := range m.Packages {
		if err := generatePackage(rootPath, pkg, opts); err != nil {
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

func writeToFile(code []byte, rootPath, zserioPkg, fn string, opts *options) (retErr error) {
	ids := strings.Split(strings.ToLower(zserioPkg), ".")
	outputDir := path.Join(append([]string{rootPath}, ids...)...)
	outputFile, err := assembleUniqueFilePath(outputDir, fn, opts.maxPathLength)
	if err != nil {
		return fmt.Errorf("file path: %w", err)
	}

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

	return writeGoSource(f, code, opts.doNotFormatSource)
}

func assembleUniqueFilePath(outputDir string, fn string, maxPathLength int) (string, error) {
	if uniquePaths == nil {
		uniquePaths = make(map[string]int)
	}

	if minLength := len(outputDir) + FileSuffixLength + MinFileNameLength; maxPathLength < minLength {
		return "", fmt.Errorf("maximum path length %d is too short, at least %d required to fit output directory", maxPathLength, minLength)
	}

	outputFile := path.Join(outputDir, fn)

	if maxPathLength > 0 {
		maxLength := maxPathLength - FileSuffixLength
		if len(outputFile) > maxLength {
			outputFile = outputFile[:maxLength]
			if strings.Contains(outputFile, "_") && strings.LastIndex(outputFile, "_") > strings.LastIndex(outputFile, "/") {
				outputFile = outputFile[:strings.LastIndex(outputFile, "_")]
			}
		}
	}

	if count, ok := uniquePaths[outputFile]; ok {
		uniquePaths[outputFile]++
		return fmt.Sprintf("%s_%d%s", outputFile, count, FileSuffix), nil
	} else {
		uniquePaths[outputFile] = 1
		return fmt.Sprintf("%s%s", outputFile, FileSuffix), nil
	}
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
	options      *options
}

func newOutput(name, template string, templateData data, withPreamble bool, opts *options) output {
	newData := data{
		"options": opts,
	}
	maps.Copy(newData, templateData)
	return output{
		name:         name,
		template:     template,
		data:         newData,
		withPreamble: withPreamble,
		options:      opts,
	}
}

func newTypeOutput(astType any, d data, withPreamble bool, opts *options) output {
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
	return newOutput(strcase.ToSnake(name), typeName+".go.tmpl", templateData, withPreamble, opts)
}

func (o *output) executeTemplate(w io.Writer) error {
	if o.withPreamble {
		if err := executeTemplate(w, "preamble.go.tmpl", o.data); err != nil {
			return fmt.Errorf("generate preamble: %w", err)
		}
	}

	return executeTemplate(w, o.template, o.data)
}

func newOutputs(pkg *ast.Package, rootPackage string, opts *options) []output {
	data := data{
		"pkg":         pkg,
		"rootPackage": rootPackage,
	}

	var outs []output

	for _, t := range pkg.Enums {
		outs = append(outs, newTypeOutput(t, data, !opts.outputToStdout, opts))
	}

	for _, t := range pkg.Unions {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !opts.outputToStdout, opts))
	}

	for _, t := range pkg.Bitmasks {
		outs = append(outs, newTypeOutput(t, data, !opts.outputToStdout, opts))
	}

	for _, t := range pkg.Choices {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !opts.outputToStdout, opts))
	}

	for _, t := range pkg.Structs {
		if len(t.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		outs = append(outs, newTypeOutput(t, data, !opts.outputToStdout, opts))
	}

	sort.Slice(outs, func(i, j int) bool {
		if outs[i].template == outs[j].template {
			return outs[i].name > outs[j].name
		}

		return outs[i].template > outs[j].template
	})

	return append([]output{newOutput("pkg", "package.go.tmpl", data, true, opts)}, outs...)
}

func generatePackage(rootPath string, pkg *ast.Package, opts *options) error {
	if opts.rootPackage == "" {
		return errors.New("BUG: root package needs to be specified")
	}

	outs := newOutputs(pkg, opts.rootPackage, opts)

	if opts.outputToStdout {
		var out bytes.Buffer
		for _, o := range outs {
			if err := o.executeTemplate(&out); err != nil {
				return fmt.Errorf("generate %s: %w", o.name, err)
			}
		}

		return writeGoSource(os.Stdout, out.Bytes(), opts.doNotFormatSource)
	}

	for _, o := range outs {
		var out bytes.Buffer
		if err := o.executeTemplate(&out); err != nil {
			return fmt.Errorf("generate %s: %w", o.name, err)
		}

		if err := writeToFile(out.Bytes(), rootPath, pkg.Name, o.name, opts); err != nil {
			return fmt.Errorf("write %s: %w", o.name, err)
		}
	}

	return nil
}
