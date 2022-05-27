package nodes

// ValueNode represents node with any type value
type ValueNode struct {
	value any
}

// NewValueNode returns new ValueNode
func NewValueNode(value any) ValueNode {
	return ValueNode{value: value}
}

// Accept accepts translate visitor to invoke TranslateValueNode method
func (n ValueNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateValueNode(n)
}

// Value returns value of node
func (n ValueNode) Value() any {
	return n.value
}
