package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/woven-planet/go-zserio/internal/ast"
)

type TypeError struct {
	ast.TypeReference
	Field string
	Err   error
}

func (e TypeError) Error() string {
	return fmt.Sprintf("%s.%s, field %s: %s", e.TypeReference.Package, e.TypeReference.Name, e.Field, e.Err.Error())
}

type UnknownPackage struct {
	Package string
}

func (e UnknownPackage) Error() string {
	return fmt.Sprintf("unknown package: %s", e.Package)
}

type UnknownType struct {
	ast.TypeReference
}

func (e UnknownType) Error() string {
	return fmt.Sprintf("unknown type: %s.%s (type not found in package)", e.TypeReference.Package, e.TypeReference.Name)
}

var ErrIncorrectNumberOfTemplateParameters = errors.New("incorrect number of template parameters")

func (m *Model) CapitalizeNames() {
	for _, pkg := range m.Packages {
		for _, e := range pkg.Enums {
			e.Name = strings.Title(e.Name)
			if !e.Type.IsBuiltin {
				e.Type.Name = strings.Title(e.Type.Name)
			}

			for _, item := range e.Items {
				item.Name = strings.Title(item.Name)
			}
		}

		for _, u := range pkg.Unions {
			u.Name = strings.Title(u.Name)
			for _, typeParameter := range u.TypeParameters {
				typeParameter.Name = strings.Title(typeParameter.Name)
			}
			for _, field := range u.Fields {
				field.Name = strings.Title(field.Name)
				if !field.Type.IsBuiltin {
					field.Type.Name = strings.Title(field.Type.Name)
				}
			}
			for _, function := range u.Functions {
				function.Name = strings.Title(function.Name)
			}
		}

		for _, bitmask := range pkg.Bitmasks {
			bitmask.Name = strings.Title(bitmask.Name)
			for _, value := range bitmask.Values {
				value.Name = strings.Title(value.Name)
			}
		}

		for _, structure := range pkg.Structs {
			for _, typeParameter := range structure.TypeParameters {
				typeParameter.Name = strings.Title(typeParameter.Name)
			}
			for _, field := range structure.Fields {
				field.Name = strings.Title(field.Name)
				if !field.Type.IsBuiltin {
					field.Type.Name = strings.Title(field.Type.Name)
				}
			}
			for _, function := range structure.Functions {
				function.Name = strings.Title(function.Name)
			}

		}

		for _, choice := range pkg.Choices {
			for _, typeParameter := range choice.TypeParameters {
				typeParameter.Name = strings.Title(typeParameter.Name)
			}
			for _, choiceCase := range choice.Cases {
				if choiceCase.Field != nil {
					choiceCase.Field.Name = strings.Title(choiceCase.Field.Name)
					if !choiceCase.Field.Type.IsBuiltin {
						choiceCase.Field.Type.Name = strings.Title(choiceCase.Field.Type.Name)
					}
				}
			}
			for _, function := range choice.Functions {
				function.Name = strings.Title(function.Name)
			}
		}
	}
}

func (m *Model) EvaluateEnums() error {
	for _, pkg := range m.Packages {
		for _, e := range pkg.Enums {
			if err := e.Evaluate(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Model) EvaluateUnions() error {
	for _, pkg := range m.Packages {
		for _, u := range pkg.Unions {
			if err := u.Evaluate(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Model) EvaluateBitmasks() error {
	for _, pkg := range m.Packages {
		for _, bitmask := range pkg.Bitmasks {
			if err := bitmask.Evaluate(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Model) EvaluateStructs() error {
	for _, pkg := range m.Packages {
		for _, s := range pkg.Structs {
			if err := s.Evaluate(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Model) EvaluateChoices() error {
	for _, pkg := range m.Packages {
		for _, choice := range pkg.Choices {
			if err := choice.Evaluate(pkg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Model) InstantiateTemplates() error {
	var err error
	for _, pkg := range m.Packages {
		for _, instantiatedType := range pkg.InstantiatedTypes {
			if instantiatedType.Type, err = m.instantiateType(
				pkg,
				instantiatedType.Type,
				instantiatedType.Type.TemplateArguments,
				instantiatedType.Name); err != nil {
				return TypeError{
					TypeReference: ast.TypeReference{
						IsBuiltin: false,
						Package:   pkg.Name,
						Name:      instantiatedType.Name,
					},
					Err: err,
				}
			}
		}

		for _, structure := range pkg.Structs {
			// don't instantiate template structs
			if len(structure.TemplateParameters) > 0 {
				continue
			}
			for _, field := range structure.Fields {
				if field.Type, err = m.instantiateType(
					pkg,
					field.Type,
					field.Type.TemplateArguments,
					""); // auto-generate a new name
				err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      structure.Name,
						},
						Field: field.Name,
						Err:   err,
					}
				}
			}
		}
	}
	return nil
}

func (m *Model) instantiateStruct(
	pkgScope *ast.Package,
	templStruct *ast.Struct,
	templateArguments []*ast.TypeReference,
	instantiatedName string) (*ast.TypeReference, error) {
	if len(templateArguments) == 0 {
		return nil, errors.New("struct instantiation called without parameters")
	}
	if len(templStruct.TemplateParameters) != len(templateArguments) {
		return nil, ErrIncorrectNumberOfTemplateParameters
	}
	// templateTypes is a map from the template parameter name to the instantiated
	// type.
	templateTypes := map[string]*ast.TypeReference{}
	for ix, ta := range templateArguments {
		// Instantiate the subtype
		ta, err := m.instantiateType(pkgScope, ta, ta.TemplateArguments, "")
		if err != nil {
			return nil, err
		}
		templateTypes[templStruct.TemplateParameters[ix]] = ta
	}

	newType := &ast.TypeReference{
		IsBuiltin: false,
		Package:   pkgScope.Name,
		Name:      instantiatedName,
	}
	if _, exists := pkgScope.Structs[newType.Name]; !exists {
		structure := &ast.Struct{
			Name:           newType.Name,
			Comment:        templStruct.Comment,
			Fields:         []*ast.Field{},
			TypeParameters: []*ast.Parameter{},
		}

		// resolve the templates for the fields
		for _, field := range templStruct.Fields {
			fieldCopy := *field
			if len(field.Type.TemplateArguments) > 0 {
				// the field itself is a templated compound type, and needs to
				// be instantiated first.

				// the template argument parameter must be recreated, as the
				// field template arguments might differ from the struct template
				// arguments. Example:
				// struct<A, B, C> {
				//    foo<int, C> field1;
				//    bar<B, A> field2;
				// }
				newTemplateParameters := append([]*ast.TypeReference{}, field.Type.TemplateArguments...)
				for ix, templateParameter := range newTemplateParameters {
					if passedTemplateArgument, found := templateTypes[templateParameter.Name]; found {
						newTemplateParameters[ix] = passedTemplateArgument
					}
				}

				tt, err := m.instantiateType(pkgScope, field.Type, newTemplateParameters, "")
				if err != nil {
					return nil, err
				}
				fieldCopy.Type = tt
				structure.Fields = append(structure.Fields, &fieldCopy)
				continue
			} else if tt := templateTypes[field.Type.Name]; tt != nil && field.Type.Package == "" {
				// the field a template itself
				fieldCopy.Type = tt
				//  If the template has type arguments, these arguments need to
				// be applied also to the instantiated type.
				fieldCopy.Type.TypeArguments = field.Type.TypeArguments
				structure.Fields = append(structure.Fields, &fieldCopy)
			} else {
				// not templated
				structure.Fields = append(structure.Fields, field)
			}
		}

		// resolve the templates for the types
		for _, typeParam := range templStruct.TypeParameters {
			if tt := templateTypes[typeParam.Type.Name]; tt != nil && typeParam.Type.Package == "" {
				// the parameter is a template
				structure.TypeParameters = append(structure.TypeParameters, &ast.Parameter{
					Name: typeParam.Name,
					Type: tt,
				})
			} else {
				// not templated
				structure.TypeParameters = append(structure.TypeParameters, typeParam)
			}
		}

		// update the scope with the newly created scope
		pkgScope.Structs[newType.Name] = structure
		pkgScope.LocalSymbols.TypeScope[newType.Name] = structure
		structure.BuildScope(pkgScope)
	}

	return newType, nil
}

func generateInstantiatedName(tr *ast.TypeReference, templateArguments []*ast.TypeReference) string {
	name := []string{strings.Title(tr.Name)}
	for _, ta := range templateArguments {
		name = append(name, strings.Title(ta.Name))
	}
	return strings.Join(name, "")
}

func (m *Model) instantiateType(pkgScope *ast.Package, tr *ast.TypeReference, templateArguments []*ast.TypeReference, instantiatedName string) (*ast.TypeReference, error) {
	if len(templateArguments) == 0 {
		return tr, nil
	}

	if instantiatedName == "" {
		// No explicit name  for the new instantiated object was given.
		// We need to create an instance of tr, using the arguments from
		// templateArguments. First we build up the name of the instantiated template
		// by appending all the argument type names.
		instantiatedName = generateInstantiatedName(tr, templateArguments)
	}

	templateSymbol, err := pkgScope.GetSymbol(tr.Name)
	if err != nil {
		return nil, err
	}
	var instantiatedType *ast.TypeReference

	switch n := templateSymbol.Symbol.(type) {
	case *ast.Struct:
		instantiatedType, err = m.instantiateStruct(pkgScope, n, templateArguments, instantiatedName)
	default:
		return nil, errors.New("Template instantiation not supported for this type")
	}

	if err != nil {
		return nil, err
	}

	// Link the argument expressions from the template type to the instantiation
	instantiatedType.TypeArguments = tr.TypeArguments
	return instantiatedType, nil
}

func (m *Model) ResolveImports() error {
	for _, pkg := range m.Packages {
		for _, imp := range pkg.Imports {
			iPkg, found := m.Packages[imp.Package]
			if !found {
				return UnknownPackage{imp.Package}
			}
			if imp.Type == "*" {
				pkg.ImportedPackages = append(pkg.ImportedPackages, iPkg)
			}
		}
	}
	return nil
}

// This assumes that Template Instantiation has already happened
func (m *Model) ResolveTypes() error {
	for _, pkg := range m.Packages {
		for _, structure := range pkg.Structs {
			templateParams := map[string]bool{}
			for _, p := range structure.TemplateParameters {
				templateParams[p] = true
			}

			for _, param := range structure.TypeParameters {
				if param.Type.Package == "" && templateParams[param.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, param.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      structure.Name,
						},
						Field: param.Name,
						Err:   err,
					}
				}
			}

			for _, field := range structure.Fields {
				if field.Type.Package == "" && templateParams[field.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, field.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      structure.Name,
						},
						Field: field.Name,
						Err:   err,
					}
				}
			}

			for _, functions := range structure.Functions {
				if err := m.resolveTypeReference(pkg, functions.ReturnType, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      structure.Name,
						},
						Field: functions.Name,
						Err:   err,
					}
				}
			}
		}

		for _, choice := range pkg.Choices {
			templateParams := map[string]bool{}
			for _, p := range choice.TemplateParameters {
				templateParams[p] = true
			}

			for _, param := range choice.TypeParameters {
				if param.Type.Package == "" && templateParams[param.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, param.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      choice.Name,
						},
						Field: param.Name,
						Err:   err,
					}
				}
			}

			for _, choiceCase := range choice.Cases {
				if choiceCase.Field == nil {
					continue
				}
				if choiceCase.Field.Type.Package == "" && templateParams[choiceCase.Field.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, choiceCase.Field.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      choice.Name,
						},
						Field: choiceCase.Field.Name,
						Err:   err,
					}
				}
			}
		}

		for _, union := range pkg.Unions {
			templateParams := map[string]bool{}
			for _, p := range union.TemplateParameters {
				templateParams[p] = true
			}

			for _, param := range union.TypeParameters {
				if param.Type.Package == "" && templateParams[param.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, param.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      union.Name,
						},
						Field: param.Name,
						Err:   err,
					}
				}
			}

			for _, field := range union.Fields {
				if field.Type.Package == "" && templateParams[field.Type.Name] {
					// Ignore template types
					continue
				}

				if err := m.resolveTypeReference(pkg, field.Type, templateParams); err != nil {
					return TypeError{
						TypeReference: ast.TypeReference{
							IsBuiltin: false,
							Package:   pkg.Name,
							Name:      union.Name,
						},
						Field: field.Name,
						Err:   err,
					}
				}
			}
		}

		for _, subtype := range pkg.Subtypes {
			if err := m.resolveTypeReference(pkg, subtype.Type, map[string]bool{}); err != nil {
				return TypeError{
					TypeReference: ast.TypeReference{
						IsBuiltin: false,
						Package:   pkg.Name,
						Name:      subtype.Type.Name,
					},
					Field: subtype.Name,
					Err:   err,
				}
			}
		}

		for _, bitmask := range pkg.Bitmasks {
			if err := m.resolveTypeReference(pkg, bitmask.Type, map[string]bool{}); err != nil {
				return TypeError{
					TypeReference: ast.TypeReference{
						IsBuiltin: false,
						Package:   pkg.Name,
						Name:      bitmask.Type.Name,
					},
					Field: bitmask.Name,
					Err:   err,
				}
			}
		}
		for _, instantiatedType := range pkg.InstantiatedTypes {
			if err := m.resolveTypeReference(pkg, instantiatedType.Type, map[string]bool{}); err != nil {
				return TypeError{
					TypeReference: ast.TypeReference{
						IsBuiltin: false,
						Package:   pkg.Name,
						Name:      instantiatedType.Type.Name,
					},
					Field: instantiatedType.Name,
					Err:   err,
				}
			}
		}
	}
	return nil
}

func (m *Model) resolveTypeReference(pkgScope *ast.Package, tr *ast.TypeReference, ignoreParameters map[string]bool) error {
	if tr.IsBuiltin {
		return nil
	}

	if tr.Package != "" {
		fPkg, found := m.Packages[tr.Package]
		if !found {
			return UnknownPackage{tr.Package}
		}
		if !fPkg.HasType(tr.Name) {
			return UnknownType{TypeReference: *tr}
		}
	} else {
		if pkgScope.HasType(tr.Name) {
			tr.Package = pkgScope.Name
		} else {
			for _, imp := range pkgScope.Imports {
				iPkg, found := m.Packages[imp.Package]
				if !found {
					return UnknownPackage{imp.Package}
				}
				if (imp.Type == "*" || imp.Type == tr.Name) && iPkg.HasType(tr.Name) {
					tr.Package = imp.Package
					break
				}
			}
		}
	}

	if tr.Package == "" {
		return UnknownType{TypeReference: *tr}
	}

	for _, ta := range tr.TemplateArguments {
		if !ignoreParameters[ta.Name] {
			if err := m.resolveTypeReference(pkgScope, ta, ignoreParameters); err != nil {
				return err
			}
		}
	}

	return nil
}
