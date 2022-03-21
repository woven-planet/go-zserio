package generator

import (
	"embed"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed templates/*.go.tmpl
var content embed.FS

var templates *template.Template

func init() {
	var err error
	funcs := map[string]any{
		"isDeltaPackable":  IsDeltaPackable,
		"goType":           GoType,
		"goPackageName":    GoPackageName,
		"goGetAllImports":  GoGetAllImports,
		"getTypeParameter": GetTypeParameter,
		"goPackageAlias":   GoPackageAlias,
		"goExpression":     GoExpression,
		"goArrayTraits":    GoArrayTraits,
		"goNativeType":     GoNativeType,
		"add":              Add,
	}

	templates, err = template.New("generator").Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap(funcs)).ParseFS(content, "templates/*.go.tmpl")
	if err != nil {
		panic(err.Error())
	}
}
