package model

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
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

func FromFilesystem(paths ...string) (*Model, error) {
	if len(paths) == 0 {
		return nil, errors.New("at least a single path must be given")
	}

	m := NewModel()

	if len(paths) > 1 {
		for _, p := range paths {
			if err := m.parseFile(p); err != nil {
				return nil, err
			}
		}
	} else {
		if err := m.walkDir(paths[0]); err != nil {
			return nil, fmt.Errorf("walk dir: %w", err)
		}
	}

	if err := m.ResolveImports(); err != nil {
		return nil, fmt.Errorf("resolve imports: %w", err)
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

func (m *Model) walkDir(root string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
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
		return m.parseFile(path)
	})
}

func (m *Model) parseFile(path string) error {
	log.Printf("Parsing file %s...\n", path)
	pkg, err := PackageFromFile(path)
	if err != nil {
		return err
	}

	if err := pkg.CollectSymbols(); err != nil {
		return err
	}

	m.Packages[pkg.Name] = pkg
	return nil
}
