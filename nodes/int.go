package nodes

// IntNode represents node with int value
type IntNode int64

// NewIntNode returns new IntNode
func NewIntNode(value int64) IntNode {
	return IntNode(value)
}

// Accept accepts translate visitor to invoke TranslateIntNode method
func (n IntNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateIntNode(n)
}
