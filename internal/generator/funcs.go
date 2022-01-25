package generator

// This file contains helper functions that can be called from templates.

import (
	"strings"

	"github.com/woven-planet/go-zserio/internal/ast"
)

func GoType(scope ast.Scope, typ *ast.TypeReference) (string, error) {
	return scope.GoType(typ)
}

func GoArrayTraits(scope ast.Scope, typ *ast.TypeReference) (string, error) {
	return scope.GoArrayTraits(typ)
}

func GoPackageName(pkg string) string {
	p := strings.Split(pkg, ".")
	return strings.ToLower(p[len(p)-1])
}

// GoGetAllImports collects a list of all imports (direct and indirect) of a
// package. This is needed to first generate a full import list, out of which
// unused items will later be erased
func GoGetAllImports(pkg *ast.Package) []string {
	imports := []string{}
	for _, importPkg := range pkg.ImportedPackages {
		imports = append(imports, importPkg.Name)
		imports = append(imports, GoGetAllImports(importPkg)...)
	}
	return imports
}
func GoPackageAlias(pkg string) string {
	return ast.GoPackageAlias(pkg)
}

func GoExpression(scope ast.Scope, expression *ast.Expression) string {
	return ExpressionToGoString(scope, expression)
}

func GoNativeType(pkg *ast.Package, typ *ast.TypeReference) (*ast.NativeZserioTypeReference, error) {
	return pkg.GetZserioNativeType(typ)
}
