package nodes

// BoolNode represents node with Bool value
type BoolNode bool

// NewBoolNode returns new BoolNode
func NewBoolNode(value bool) BoolNode {
	return BoolNode(value)
}

// Accept accepts translate visitor to invoke TranslateBoolNode method
func (n BoolNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateBoolNode(n)
}

// Value returns value of the node as primitive type
func (n BoolNode) Value() interface{} {
	return bool(n)
}
