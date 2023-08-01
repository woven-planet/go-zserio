package ast

import (
	"errors"
	"fmt"
	"strings"
)

type Package struct {
	BaseScope

	Name    string
	Comment string

	Imports           []*Import
	Consts            map[string]*Const
	Subtypes          map[string]*Subtype
	InstantiatedTypes map[string]*InstantiateType
	Enums             map[string]*Enum
	Unions            map[string]*Union
	Bitmasks          map[string]*BitmaskType
	Structs           map[string]*Struct
	Choices           map[string]*Choice

	ImportedPackages []*Package
	LocalSymbols     *SymbolScope
}

// OriginalTypeReference is a reference to the original zserio type, for example
// fundamental types, but also enums/structs, that were aliased by a subtype.
type OriginalTypeReference struct {
	Type         *TypeReference
	RequiresCast bool
	IsMarshaler  bool
}

func NewPackage(name string) *Package {
	return &Package{
		BaseScope: BaseScope{
			Parent: rootScope,
		},
		Name:              name,
		Consts:            map[string]*Const{},
		Subtypes:          map[string]*Subtype{},
		InstantiatedTypes: map[string]*InstantiateType{},
		Structs:           map[string]*Struct{},
		Enums:             map[string]*Enum{},
		Unions:            map[string]*Union{},
		Bitmasks:          map[string]*BitmaskType{},
		Choices:           map[string]*Choice{},
	}
}

// HasType checks if a type is defined *in this scope*. It does not check parent
// scopes.
func (p *Package) HasType(name string) bool {
	_, err := p.LocalSymbols.GetType(name)
	return err == nil
}

// CollectSymbols collects all symbols from a package, and organizes them in the
// scope
func (p *Package) CollectSymbols() error {
	p.LocalSymbols = &SymbolScope{
		TypeScope:      make(map[string]any),
		OtherScope:     make(map[string]any),
		CompoundScopes: make(map[string]*CompoundScope),
	}
	for name, value := range p.Subtypes {
		p.LocalSymbols.TypeScope[name] = value
	}
	for name, value := range p.InstantiatedTypes {
		p.LocalSymbols.TypeScope[name] = value
	}
	for name, value := range p.Consts {
		p.LocalSymbols.TypeScope[name] = value
	}
	for name, value := range p.Structs {
		p.LocalSymbols.TypeScope[name] = value
		if err := value.BuildScope(p); err != nil {
			return err
		}
	}
	for name, value := range p.Enums {
		p.LocalSymbols.TypeScope[name] = value

		// add all enum values as symbols
		for _, enumItem := range value.Items {
			p.LocalSymbols.OtherScope[enumItem.Name] = enumItem
		}
	}
	for name, value := range p.Unions {
		p.LocalSymbols.TypeScope[name] = value
		if err := value.BuildScope(p); err != nil {
			return err
		}
	}
	for name, value := range p.Bitmasks {
		p.LocalSymbols.TypeScope[name] = value
	}
	for name, value := range p.Choices {
		p.LocalSymbols.TypeScope[name] = value
		if err := value.BuildScope(p); err != nil {
			return err
		}
	}
	return nil
}

// GoType looks up a name in the scope, and provides the Go type name for it.
// If the type is not found ErrUnkownType is returned.
func (p *Package) GoType(t *TypeReference) (string, error) {
	if !t.IsBuiltin {
		if t.Package == "" {
			return "", errors.New("type is not resolved")
		}
		if t.Package == p.Name {
			return strings.Title(t.Name), nil
		}
		parts := strings.Split(t.Package, ".")
		if len(parts) == 0 {
			return strings.Title(t.Name), nil
		}

		packageAlias := GoPackageAlias(t.Package)
		return fmt.Sprintf("%s.%s", packageAlias, strings.Title(t.Name)), nil
	}

	return p.BaseScope.GoType(t)
}

// GoType looks up a name in the scope, and provides the Go type name for it.
// If the type is not found ErrUnkownType is returned.
func (p *Package) OriginalType(t *TypeReference) (*OriginalTypeReference, error) {
	return p.GetOriginalType(t)
}

// GoArrayTraits returns the array traits object for non-basic zserio types,
// such as enums, subtypes, or structures.
func (p *Package) GoArrayTraits(t *TypeReference) (string, error) {
	if t.IsBuiltin {
		return p.BaseScope.GoArrayTraits(t)
	}

	typeScope, err := p.GetImportedScope(t.Package)
	if err != nil {
		return "", err
	}
	typeSymbol, err := typeScope.GetSymbol(t.Name)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
	}

	objArrayTraitsStr := "ztype.ObjectArrayTraits"
	switch n := typeSymbol.Symbol.(type) {
	case *Const:
		return p.GoArrayTraits(n.Type)
	case *Subtype:
		return p.GoArrayTraits(n.Type)
	case *InstantiateType, *Struct, *Enum, *Union, *BitmaskType, *Choice:
		return objArrayTraitsStr, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
	}
}

// IsDeltaPackable indicates if a non-basic zserio type (struct, enum, subtype, ...)
// can be delta-packed.
func (p *Package) IsDeltaPackable(t *TypeReference) (bool, error) {
	if t.IsBuiltin {
		return p.BaseScope.Parent.IsDeltaPackable(t)
	}

	typeScope, err := p.GetImportedScope(t.Package)
	if err != nil {
		return false, err
	}
	typeSymbol, err := typeScope.GetSymbol(t.Name)
	if err != nil {
		return false, fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
	}

	switch n := typeSymbol.Symbol.(type) {
	case *Const:
		return p.IsDeltaPackable(n.Type)
	case *Subtype:
		return p.IsDeltaPackable(n.Type)
	case *InstantiateType:
		return p.IsDeltaPackable(n.Type)
	case *Struct, *Union, *Choice:
		// compound types have their own packed marshaling function
		return false, nil
	case *Enum, *BitmaskType:
		return true, nil
	default:
		return false, fmt.Errorf("%w: %s", ErrUnknownType, t.Name)
	}
}

type SymbolReference struct {
	Symbol       any
	Name         string
	Package      string
	CompoundName string
}

// GetSymbol searches for a symbol not only in the current scope, but also in
// all imported scopes.
func (p *Package) GetSymbol(name string) (*SymbolReference, error) {
	// The search order is important to avoid any recursive lookup, in case
	// field and type have the same name.
	// first search for types in local and all imported packages
	if typeSymbol, err := p.GetType(name); err == nil {
		return typeSymbol, nil
	}

	// second, search for composite types values (struct fields), in the
	// active local scopes
	if typeSymbol, err := p.GetCurrentCompoundType(name); err == nil {
		return typeSymbol, nil
	}

	// third, search all other types (enum fields)
	if typeSymbol, err := p.GetOtherType(name); err == nil {
		return typeSymbol, nil
	}
	return nil, errors.New("symbol not found")
}

func (p *Package) GetType(name string) (*SymbolReference, error) {
	// First check the local scope
	if typeSymbol, err := p.LocalSymbols.GetType(name); err == nil {
		typeSymbol.Package = p.Name
		return typeSymbol, nil
	}

	// If the symbol was not found in the local scope, check all imported scopes
	for _, importedPackage := range p.ImportedPackages {
		if typeSymbol, err := importedPackage.LocalSymbols.GetType(name); err == nil {
			typeSymbol.Package = importedPackage.Name
			return typeSymbol, nil
		}
	}

	// Check in the scope of imported imported
	return nil, errors.New("type not found")
}

func (p *Package) GetCurrentCompoundType(name string) (*SymbolReference, error) {
	if p.LocalSymbols.CurrentCompoundScope != nil {
		if typeSymbol, err := p.LocalSymbols.GetCompoundType(*p.LocalSymbols.CurrentCompoundScope, name); err == nil {
			typeSymbol.Package = p.Name
			return typeSymbol, nil
		}
	}

	// check all imported scopes
	for _, importedPackage := range p.ImportedPackages {
		if importedPackage.LocalSymbols.CurrentCompoundScope != nil {
			if typeSymbol, err := importedPackage.LocalSymbols.GetCompoundType(*importedPackage.LocalSymbols.CurrentCompoundScope, name); err == nil {
				typeSymbol.Package = importedPackage.Name
				return typeSymbol, nil
			}
		}
	}
	return nil, errors.New("compound symbol not found")
}

func (p *Package) GetCompoundType(compoundSymbolName, name string) (*SymbolReference, error) {
	if typeSymbol, err := p.LocalSymbols.GetCompoundType(compoundSymbolName, name); err == nil {
		typeSymbol.Package = p.Name
		return typeSymbol, nil
	}

	// check all imported scopes
	for _, importedPackage := range p.ImportedPackages {
		if typeSymbol, err := importedPackage.LocalSymbols.GetCompoundType(compoundSymbolName, name); err == nil {
			typeSymbol.Package = importedPackage.Name
			return typeSymbol, nil
		}
	}
	return nil, errors.New("compound symbol not found")
}

func (p *Package) GetOtherType(name string) (*SymbolReference, error) {
	if symbolRef, err := p.LocalSymbols.GetOtherType(name); err == nil {
		symbolRef.Package = p.Name
		return symbolRef, nil
	}

	// check all imported scopes
	for _, importedPackage := range p.ImportedPackages {
		if symbolRef, err := importedPackage.GetOtherType(name); err == nil {
			return symbolRef, nil
		}
	}
	return nil, errors.New("other type not found")
}

func (p *Package) GetOriginalSymbol(symbolReference *SymbolReference) (*SymbolReference, error) {
	var typeRef *TypeReference
	var err error
	counter := 0
	for {
		switch n := symbolReference.Symbol.(type) {
		case *Enum:
			return symbolReference, nil
		case *EnumItem:
			return symbolReference, nil
		case *Union:
			return symbolReference, nil
		case *Struct:
			return symbolReference, nil
		case *Choice:
			return symbolReference, nil
		case *Field:
			typeRef = n.Type
		case *Parameter:
			typeRef = n.Type
		case *Subtype:
			typeRef = n.Type
		case *TypeReference:
			typeRef = n
		case *Function:
			return symbolReference, nil
		case *BitmaskType:
			return symbolReference, nil
		default:
			return nil, errors.New("unable to evaluate the expression type")
		}
		if typeRef.IsBuiltin {
			return &SymbolReference{
				Symbol: typeRef,
			}, nil
		}
		p, err = p.GetImportedScope(typeRef.Package)
		if err != nil {
			return nil, err
		}
		symbolReference, err = p.GetSymbol(typeRef.Name)
		if err != nil {
			return nil, err
		}
		counter += 1
		if counter > 100 {
			return nil, errors.New("symbol lookup failed")
		}
	}

}

func (p *Package) GetOriginalType(typeRef *TypeReference) (*OriginalTypeReference, error) {
	requiresCast := false
	counter := 0
	currentScope := p
	for {
		if typeRef.IsBuiltin {
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
			}, nil
		}
		currentScope, err := currentScope.GetImportedScope(typeRef.Package)
		if err != nil {
			return nil, err
		}
		symbol, err := currentScope.GetSymbol(typeRef.Name)
		if err != nil {
			return nil, err
		}

		switch n := symbol.Symbol.(type) {
		case *Enum:
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
				IsMarshaler:  true,
			}, nil
		case *Union:
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
				IsMarshaler:  true,
			}, nil
		case *BitmaskType:
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
				IsMarshaler:  true,
			}, nil
		case *Choice:
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
				IsMarshaler:  true,
			}, nil
		case *Struct:
			return &OriginalTypeReference{
				Type:         typeRef,
				RequiresCast: requiresCast,
				IsMarshaler:  true,
			}, nil
		case *Field:
			typeRef = n.Type
		case *Parameter:
			typeRef = n.Type
		case *Subtype:
			typeRef = n.Type
		default:
			return nil, errors.New("unable to find the native zserio type")
		}
		requiresCast = true
		counter++
		if counter > 100 {
			return nil, errors.New("internal failure to resolve type (circular lookup)")

		}
	}
}

func GoPackageAlias(packageName string) string {
	return strings.Replace(packageName, ".", "_", -1)
}

// GetTypeParameter returns the parameters of a referenced type
func (p *Package) GetTypeParameter(typeRef *TypeReference) ([]*Parameter, error) {
	nativeType, err := p.GetOriginalType(typeRef)
	if err != nil {
		return nil, err
	}
	typeScope, err := p.GetImportedScope(nativeType.Type.Package)
	if err != nil {
		return nil, err
	}
	symbol, err := typeScope.GetSymbol(nativeType.Type.Name)
	if err != nil {
		return nil, err
	}

	var parameters []*Parameter
	switch n := symbol.Symbol.(type) {
	case *Union:
		parameters = n.TypeParameters
	case *Choice:
		parameters = n.TypeParameters
	case *Struct:
		parameters = n.TypeParameters
	default:
		return nil, errors.New("type reference is not parameterizable")
	}

	if len(parameters) != len(typeRef.TypeArguments) {
		return nil, fmt.Errorf("expected %d type parameters, got %d", len(typeRef.TypeArguments), len(parameters))
	}
	return parameters, nil
}

// GetImportedScope returns a scope, if it can be found within the imported scope
func (p *Package) GetImportedScope(name string) (*Package, error) {
	if name == "" {
		return nil, errors.New("unresolved package scope")
	}
	if name == p.Name {
		return p, nil
	}
	for _, pkg := range p.ImportedPackages {
		if pkg.Name == name {
			return pkg, nil
		}
	}
	// still not found? try the imports of all imports
	for _, pkg := range p.ImportedPackages {
		obj, err := pkg.GetImportedScope(name)
		if err == nil {
			return obj, nil
		}
	}
	return nil, errors.New("scope does not exist, or has not been imported")
}
