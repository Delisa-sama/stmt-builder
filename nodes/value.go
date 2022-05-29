package nodes

// ValueNode represents node with any type value
type ValueNode struct {
	value interface{}
}

// NewValueNode returns new ValueNode
func NewValueNode(value interface{}) ValueNode {
	return ValueNode{value: value}
}

// Accept accepts translate visitor to invoke TranslateValueNode method
func (n ValueNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateValueNode(n)
}

// Value returns value of node
func (n ValueNode) Value() interface{} {
	return n.value
}
