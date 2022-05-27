package nodes

// ArrayNode represents node with multiple childs
type ArrayNode struct {
	value []Node
}

// NewArrayNode returns new ArrayNode
func NewArrayNode(value []Node) ArrayNode {
	return ArrayNode{
		value: value,
	}
}

// Childs returns collection of childs
func (n ArrayNode) Childs() []Node {
	return n.value
}

// Accept accepts translate visitor to invoke TranslateArrayNode method
func (n ArrayNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateArrayNode(n)
}
