package model

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/woven-planet/go-zserio/internal/ast"
)

type Model struct {
	Packages map[string]*ast.Package
}

func NewModel() *Model {
	return &Model{
		Packages: map[string]*ast.Package{},
	}
}

func FromFilesystem(root string) (*Model, error) {
	m := NewModel()

	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") ||
			!(strings.HasSuffix(d.Name(), ".zs") || strings.HasSuffix(d.Name(), ".zserio")) {
			return nil
		}

		pkg, err := PackageFromFile(path)
		if err != nil {
			return err
		}
		err = pkg.CollectSymbols()
		if err != nil {
			return err
		}
		m.Packages[pkg.Name] = pkg
		return nil
	}); err != nil {
		return nil, err
	}

	if err := m.ResolveImports(); err != nil {
		return nil, err
	}

	if err := m.ResolveTypes(); err != nil {
		return nil, err
	}

	if err := m.InstantiateTemplates(); err != nil {
		return nil, err
	}

	if err := m.EvaluateEnums(); err != nil {
		return nil, err
	}

	if err := m.EvaluateUnions(); err != nil {
		return nil, err
	}

	if err := m.EvaluateBitmasks(); err != nil {
		return nil, err
	}

	if err := m.EvaluateStructs(); err != nil {
		return nil, err
	}

	if err := m.EvaluateChoices(); err != nil {
		return nil, err
	}

	m.CapitalizeNames()

	return m, nil
}
