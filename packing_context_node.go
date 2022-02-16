package zserio

type IDeltaContext interface {
	ReadDescriptor(r Reader) error
	WriteDescriptor(w Writer) error
	BitSizeOfDescriptor() int
}

// PackingContextNode is a context for writing packed data.
type PackingContextNode struct {
	children []*PackingContextNode

	// the delta context used in a packed context.
	Context IDeltaContext
}

// AddChild adds a new child to a PackingContextNode.
func (context *PackingContextNode) AddChild(child *PackingContextNode) {
	context.children = append(context.children, child)
}

// GetChildren returns the child PackingContextNodes of a PackingContextNode.
func (context *PackingContextNode) GetChildren() []*PackingContextNode {
	return context.children
}

// HasContext reurns true if the packing context has a delta context.
func (context *PackingContextNode) HasContext() bool {
	return (context.Context != nil)
}

// ReadDescriptor returns the delta context of the packing context.
func (context *PackingContextNode) ReadDescriptor(r Reader) error {
	return context.Context.ReadDescriptor(r)
}

func (context *PackingContextNode) WriteDescriptor(w Writer) error {
	return context.Context.WriteDescriptor(w)
}

func (context *PackingContextNode) BitSizeOfDescriptor() int {
	return context.Context.BitSizeOfDescriptor()
}
