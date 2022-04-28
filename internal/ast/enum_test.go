package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/woven-planet/go-zserio/internal/ast"
)

func TestEnumGoName(t *testing.T) {
	tests := map[string]struct {
		item *ast.EnumItem
		want string
	}{
		"basic-uppercase": {
			item: (&ast.Enum{Name: "Topping"}).AddItem(&ast.EnumItem{Name: "PINEAPPLE"}),
			want: "ToppingPineapple",
		},
		"snake-uppercase": {
			item: (&ast.Enum{Name: "Topping"}).AddItem(&ast.EnumItem{Name: "GREEN_PEPPER"}),
			want: "ToppingGreenPepper",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.item.GoName())
		})
	}
}
