package generator

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

type Options struct {
	RootPackage       string
	DoNotFormatSource bool
}

func Generate(m *model.Model, rootPath string, options *Options) error {
	for _, pkg := range m.Packages {
		if err := generatePackage(rootPath, pkg, options); err != nil {
			return err
		}
	}
	return nil
}

func writeSource(rootPath, zserioPkg, fn string, doNotFormatCode bool, tmpl string, data map[string]interface{}) error {
	ids := strings.Split(strings.ToLower(zserioPkg), ".")
	outputDir := path.Join(append([]string{rootPath}, ids...)...)
	outputFile := path.Join(outputDir, fn)

	var out bytes.Buffer
	var err error

	templ := templates.Lookup(tmpl)
	if templ == nil {
		panic("Package template is missing")
	}

	if err = templ.Execute(&out, data); err != nil {
		return err
	}

	code := out.Bytes()

	if !doNotFormatCode {
		code, err = StripImports(out.Bytes())
		if err != nil {
			return err
		}
		code, err = format.Source(code)
		if err != nil {
			return err
		}
	}

	log.Printf("Writing %s\n", outputFile)
	if err = os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(outputFile, FileOpenFlags, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(code); err != nil {
		_ = os.Remove(outputFile)
		return err
	}
	return nil
}

func generatePackage(rootPath string, pkg *ast.Package, options *Options) error {
	var err error

	if err = writeSource(rootPath, pkg.Name, "pkg.go", options.DoNotFormatSource, "package.go.tmpl", map[string]interface{}{
		"pkg":         pkg,
		"rootPackage": options.RootPackage,
	}); err != nil {
		return err
	}

	for _, e := range pkg.Enums {
		if err = writeSource(rootPath, pkg.Name, strcase.ToSnake(e.Name)+".go", options.DoNotFormatSource, "enum.go.tmpl", map[string]interface{}{
			"pkg":         pkg,
			"enum":        e,
			"rootPackage": options.RootPackage,
		}); err != nil {
			return err
		}
	}

	for _, u := range pkg.Unions {
		if len(u.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		if err = writeSource(rootPath, pkg.Name, strcase.ToSnake(u.Name)+".go", options.DoNotFormatSource, "union.go.tmpl", map[string]interface{}{
			"pkg":         pkg,
			"union":       u,
			"rootPackage": options.RootPackage,
		}); err != nil {
			return err
		}
	}

	for _, b := range pkg.Bitmasks {
		if err = writeSource(rootPath, pkg.Name, strcase.ToSnake(b.Name)+".go", options.DoNotFormatSource, "bitmask.go.tmpl", map[string]interface{}{
			"pkg":         pkg,
			"bitmask":     b,
			"rootPackage": options.RootPackage,
		}); err != nil {
			return err
		}
	}

	for _, c := range pkg.Choices {
		if len(c.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		if err = writeSource(rootPath, pkg.Name, strcase.ToSnake(c.Name)+".go", options.DoNotFormatSource, "choice.go.tmpl", map[string]interface{}{
			"pkg":         pkg,
			"choice":      c,
			"rootPackage": options.RootPackage,
		}); err != nil {
			return err
		}
	}

	for _, str := range pkg.Structs {
		if len(str.TemplateParameters) > 0 {
			// We only need to generate source for instantiated templates
			continue
		}
		if err = writeSource(rootPath, pkg.Name, strcase.ToSnake(str.Name)+".go", options.DoNotFormatSource, "struct.go.tmpl", map[string]interface{}{
			"pkg":         pkg,
			"struct":      str,
			"rootPackage": options.RootPackage,
		}); err != nil {
			return err
		}
	}

	return nil
}
