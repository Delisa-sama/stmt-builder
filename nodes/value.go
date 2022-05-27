package nodes

type ValueNode struct {
	value any
}

func NewValueNode(value any) ValueNode {
	return ValueNode{value: value}
}

func (n ValueNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateValueNode(n)
}

func (n ValueNode) Value() any {
	return n.value
}
