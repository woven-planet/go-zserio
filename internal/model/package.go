package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/parser"
	"github.com/woven-planet/go-zserio/internal/visitor"
)

func PackageFromFile(path string) (*ast.Package, error) {
	input, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, err
	}

	lexer := parser.NewZserioLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewZserioParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.PackageDeclaration()
	visitor := &visitor.Visitor{}

	switch v := visitor.Visit(tree).(type) {
	case error:
		return nil, err
	case *ast.Package:
		return v, nil
	}

	panic("unexpected result from visitor")
}
