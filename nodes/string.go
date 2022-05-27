package nodes

type StringNode struct {
	value string
}

func NewStringNode(value string) StringNode {
	return StringNode{value: value}
}

func (n StringNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateStringNode(n)
}

func (n StringNode) String() string {
	return n.value
}
