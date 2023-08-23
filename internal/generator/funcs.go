package generator

// This file contains helper functions that can be called from templates.

import (
	"strings"

	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
)

func IsDeltaPackable(scope ast.Scope, typ *ast.TypeReference) (bool, error) {
	return scope.IsDeltaPackable(typ)
}

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
	packagesToProcess := []*ast.Package{}
	alreadyProcessed := map[string]bool{}

	for {
		for _, importPkg := range pkg.ImportedPackages {
			// check if this package is already processed
			if _, exists := alreadyProcessed[importPkg.Name]; exists {
				continue
			}
			alreadyProcessed[importPkg.Name] = true
			imports = append(imports, importPkg.Name)
			packagesToProcess = append(packagesToProcess, importPkg)
		}
		if len(packagesToProcess) == 0 {
			break
		}
		// pop the last entry
		pkg, packagesToProcess = packagesToProcess[len(packagesToProcess)-1], packagesToProcess[:len(packagesToProcess)-1]
	}
	return imports
}

func GetTypeParameter(pkg *ast.Package, typ *ast.TypeReference) ([]*ast.Parameter, error) {
	return pkg.GetTypeParameter(typ)
}

func GoPackageAlias(pkg string) string {
	return ast.GoPackageAlias(pkg)
}

func GoExpression(scope ast.Scope, expression *ast.Expression) string {
	return ExpressionToGoString(scope, expression)
}

func GoNativeType(pkg *ast.Package, typ *ast.TypeReference) (*ast.OriginalTypeReference, error) {
	return pkg.GetOriginalType(typ)
}

func Add(op1, op2 int) int {
	return op1 + op2
}

func hasIndexOperator(expression *ast.Expression) bool {
	if expression == nil {
		return false
	}

	if hasIndexOperator(expression.Operand1) {
		return true
	}
	if hasIndexOperator(expression.Operand2) {
		return true
	}
	if hasIndexOperator(expression.Operand3) {
		return true
	}

	if expression.Type == parser.ZserioParserINDEX {
		return true
	}
	return false
}

func HasIndexOperators(typeRef *ast.TypeReference) bool {
	for _, typeArgument := range typeRef.TypeArguments {
		if hasIndexOperator(typeArgument) {
			return true
		}
	}
	return false
}
