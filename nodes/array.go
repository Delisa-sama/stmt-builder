package nodes

type ArrayNode struct {
	value []Node
}

func NewArrayNode(value []Node) ArrayNode {
	return ArrayNode{
		value: value,
	}
}

func (n ArrayNode) Childs() []Node {
	return n.value
}

func (n ArrayNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateArrayNode(n)
}
