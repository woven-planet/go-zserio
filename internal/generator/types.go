package generator

var ZserioTypeToGoType = map[string]string{
	// integer types
	"int8":   "int8",
	"int16":  "int16",
	"int32":  "int32",
	"int64":  "int64",
	"uint8":  "uint8",
	"uint16": "uint16",
	"uint32": "uint32",
	"uint64": "uint64",
	// varint types
	"varint":    "int",
	"varint16":  "int16",
	"varint32":  "int32",
	"varint64":  "int64",
	"varuint":   "uint",
	"varsize":   "uint",
	"varuint16": "uint16",
	"varuint32": "uint32",
	"varuint64": "uint64",
	// bool types
	"bool": "bool",
	// string types
	"string": "string",
	// float types
	"float16": "float32",
	"float32": "float32",
	"float64": "float64",
	// arbitrary type
	"extern": "any",
	"bytes":  "any",
}
