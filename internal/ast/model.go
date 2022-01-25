package ast

type Import struct {
	Package string
	Type    string
}

type Const struct {
	Name            string
	Comment         string
	Type            *TypeReference
	ValueExpression *Expression
}

type Subtype struct {
	Name    string
	Comment string
	Type    *TypeReference
}

// InstantiateType is an instantiation of a template.
type InstantiateType struct {
	Name    string
	Comment string
	Type    *TypeReference
}

type Function struct {
	Name       string
	Comment    string
	Result     *Expression
	ReturnType *TypeReference
}

type Parameter struct {
	Name string
	Type *TypeReference
}
