package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/woven-planet/go-zserio/internal/ast"
	"github.com/woven-planet/go-zserio/internal/model"
)

func TestTemplateInstantiationStruct(t *testing.T) {
	tests := map[string]struct {
		testPackage           *ast.Package
		testTypeInstantiation *ast.InstantiateType
		resultingType         *ast.Struct
	}{
		"test-struct-template-instantiation-with-native-types": {
			testPackage: &ast.Package{
				Name: "test_package",
				Structs: map[string]*ast.Struct{
					"template_struct": {
						Name:               "template_struct",
						TemplateParameters: []string{"T", "V"},
						Fields: []*ast.Field{
							{
								Name: "field_a",
								Type: &ast.TypeReference{
									Name: "T",
								},
							},
							{
								Name: "field_b",
								Type: &ast.TypeReference{
									Name: "V",
								},
							},
						},
					},
				},
			},
			testTypeInstantiation: &ast.InstantiateType{
				Name: "struct_instantiation",
				Type: &ast.TypeReference{
					Name:    "template_struct",
					Package: "test_package",
					TemplateArguments: []*ast.TypeReference{
						{
							Name:      "int64",
							IsBuiltin: true,
						},
						{
							Name:      "int32",
							IsBuiltin: true,
						},
					},
				},
			},
			resultingType: &ast.Struct{
				Name: "struct_instantiation",
				Fields: []*ast.Field{
					{
						Name: "field_a",
						Type: &ast.TypeReference{
							Name:      "int64",
							IsBuiltin: true,
						},
					},
					{
						Name: "field_b",
						Type: &ast.TypeReference{
							Name:      "int32",
							IsBuiltin: true,
						},
					},
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// use the test data to build up a model
			test.testPackage.InstantiatedTypes = map[string]*ast.InstantiateType{
				test.testTypeInstantiation.Name: test.testTypeInstantiation,
			}
			m := &model.Model{
				Packages: map[string]*ast.Package{
					test.testPackage.Name: test.testPackage,
				},
			}
			if err := test.testPackage.CollectSymbols(); err != nil {
				t.Fatalf("error during symbol collection: %v", err)
			}

			if err := m.InstantiateTemplates(); err != nil {
				t.Fatalf("error during template instantiation: %v", err)
			}

			// search for the instantiated struct
			if s, ok := test.testPackage.Structs[test.resultingType.Name]; !ok {
				t.Fatalf("instantiated struct not found : %s", test.resultingType.Name)
			} else {
				if diff := cmp.Diff(test.resultingType, s); diff != "" {
					t.Errorf("template instantiation does not match: %s", diff)
				}
			}
		})
	}
}
