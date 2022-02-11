package ztype

import "github.com/woven-planet/go-zserio"

// CreatePackingContextNode creates a new PackingContextNode using a delta context.
func CreatePackingContextNode[T any]() *zserio.PackingContextNode {
	return &zserio.PackingContextNode{
		Context: &DeltaContext[T]{},
	}
}
