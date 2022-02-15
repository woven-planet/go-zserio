package reference

import (
	"os"
	"path"
	"testing"

	"gen/github.com/woven-planet/go-zserio/test/go/reference_modules/core/instantiations"
	"gen/github.com/woven-planet/go-zserio/test/go/reference_modules/core/types"
	"gen/github.com/woven-planet/go-zserio/test/go/reference_modules/testobject1/testobject"
	zserio "github.com/woven-planet/go-zserio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testWorkspace(filePath string) string {
	return path.Join(os.Getenv("TEST_SRCDIR"), os.Getenv("TEST_WORKSPACE"), filePath)
}

var (
	// ReferenceFilePath is the path to the input data.
	ReferenceFilePath string = testWorkspace(os.Getenv("TESTDATA_BIN"))
)

func TestRoundtripFromZserio(t *testing.T) {
	// Given
	want, err := os.ReadFile(ReferenceFilePath)
	require.NoError(t, err)

	// When
	var testObject testobject.TestObject
	require.NoError(t, zserio.Unmarshal(want, &testObject), "unmarshal")
	got, err := zserio.Marshal(&testObject)
	require.NoError(t, err, "marshal")

	// Then
	assert.Equal(t, want, got)
}

func TestEqualValues(t *testing.T) {
	// Given
	bytes, err := os.ReadFile(ReferenceFilePath)
	require.NoError(t, err)

	// When
	var got testobject.TestObject
	require.NoError(t, zserio.Unmarshal(bytes, &got))

	// Then
	assert.Equal(t, want(), testObject)
}

func want() testobject.TestObject {
	var d testobject.TestObject
	d.Parameter1 = 7
	d.Struct1 = instantiations.InstantiatedTemplateStruct{
		Parameter: d.Parameter1,
		Field: types.ValueWrapper{
			Parameter:   d.Parameter1,
			Value:       72,
			OtherValue:  121,
			EnumValue:   types.ColorBLUE,
			Description: "test data",
		},
	}
	d.Color1 = types.ColorRED

	d.Parameter2 = 12
	for i := 0; i < 22; i++ {
		d.StructArray = append(d.StructArray, instantiations.InstantiatedTemplateStruct{
			Parameter: d.Parameter2,
			Field: types.ValueWrapper{
				Parameter:   d.Parameter2,
				Value:       int32(i + 100),
				OtherValue:  int8(i),
				Description: "some dummy description",
				// FIXME @aignas 2022-02-15: is this a bug? How can we reserialize this back to the correct binary but not have this field set?
				// EnumValue:   types.ColorBLACK,
			},
		})
	}

	for i := 0; i < 100; i++ {
		d.BitmaskArray = append(d.BitmaskArray, types.CityAttributes(2))
	}

	// d.OptionChoice1 is deliberately not set
	d.Parameter3 = 12
	d.OptionChoice2 = &instantiations.InstantiatedTemplateChoice{
		Parameter:    uint8(d.Parameter3), // this conversion is a bit fishy
		DefaultField: 707,
	}

	// The next test is for the correct lookup of choice selector types. The
	// choice below is using an enum, whose values also exist in a different enum
	d.ChoiceSelector = types.SomeEnumHAS_A
	d.BasicChoice = types.BasicChoice{
		Type: d.ChoiceSelector,
		HasA: 5,
	}

	d.Foo = 42
	return d
}
