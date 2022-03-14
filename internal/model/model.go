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

func FromFiles(paths ...string) (*Model, error) {
	m := NewModel()

	for _, p := range paths {
		if err := m.parseFile(p); err != nil {
			return nil, err
		}
	}

	if err := m.evaluate(); err != nil {
		return nil, err
	}

	return m, nil
}

func FromFilesystem(paths ...string) (*Model, error) {
	if len(paths) == 0 {
		return nil, errors.New("at least a single path must be given")
	}

	if len(paths) != 1 {
		return FromFiles(paths...)
	}

	paths, err := walkDir(paths[0])
	if err != nil {
		return nil, err
	}

	return FromFiles(paths...)
}

func walkDir(root string) ([]string, error) {
	var paths []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		if strings.HasSuffix(d.Name(), ".zs") || strings.HasSuffix(d.Name(), ".zserio") {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walkdir: %w", err)
	}

	return paths, nil
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

func (m *Model) evaluate() error {
	if err := m.ResolveImports(); err != nil {
		return fmt.Errorf("resolve imports: %w", err)
	}

	if err := m.ResolveTypes(); err != nil {
		return fmt.Errorf("resolve types: %w", err)
	}

	if err := m.InstantiateTemplates(); err != nil {
		return fmt.Errorf("instantiate templates: %w", err)
	}

	if err := m.EvaluateConsts(); err != nil {
		return fmt.Errorf("eval consts: %w", err)
	}

	if err := m.EvaluateEnums(); err != nil {
		return fmt.Errorf("eval enums: %w", err)
	}

	if err := m.EvaluateUnions(); err != nil {
		return fmt.Errorf("eval unions: %w", err)
	}

	if err := m.EvaluateBitmasks(); err != nil {
		return fmt.Errorf("eval bitmasks: %w", err)
	}

	if err := m.EvaluateStructs(); err != nil {
		return fmt.Errorf("eval structs: %w", err)
	}

	if err := m.EvaluateChoices(); err != nil {
		return fmt.Errorf("eval choices: %w", err)
	}

	m.CapitalizeNames()
	return nil
}
