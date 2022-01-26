package zserio

import (
	"github.com/icza/bitio"
)

type IDeltaContext interface {
	ReadDescriptor(r *bitio.CountReader) error
	WriteDescriptor(w *bitio.CountWriter) error
	BitSizeOfDescriptor() int
}

// PackingContextNode is a context for writing packed data.
type PackingContextNode struct {
	children []*PackingContextNode

	// the delta context used in a packed context.
	Context IDeltaContext
}

// CreateChild adds a new child to a PackingContextNode.
func (context *PackingContextNode) AddChild(child *PackingContextNode) {
	context.children = append(context.children, child)
}

// GetChildren returns the child PackingContextNodes of a PackingContextNode.
func (context *PackingContextNode) GetChildren() []*PackingContextNode {
	return context.children
}

// HasDeltaContext reurns true if the packing context has a delta context.
func (context *PackingContextNode) HasContext() bool {
	return (context.Context != nil)
}

// GetDeltaContext returns the delta context of the packing context.
func (packingNode *PackingContextNode) ReadDescriptor(r *bitio.CountReader) error {
	return packingNode.Context.ReadDescriptor(r)
}

func (packingNode *PackingContextNode) WriteDescriptor(w *bitio.CountWriter) error {
	return packingNode.Context.WriteDescriptor(w)
}

func (packingNode *PackingContextNode) BitSizeOfDescriptor() int {
	return packingNode.Context.BitSizeOfDescriptor()
}
