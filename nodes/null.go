package nodes

// NullNode represents node with nil value
type NullNode struct{}

// NewNullNode returns new NullNode
func NewNullNode() NullNode {
	return NullNode{}
}

// Accept accepts translate visitor to invoke TranslateNullNode method
func (n NullNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNullNode(n)
}

// Value returns value of the node as primitive type
func (n NullNode) Value() interface{} {
	return nil
}
