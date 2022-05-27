package nodes

// OrNode represents OR operator
type OrNode struct {
	left  Node
	right Node
}

// NewOrNode returns new OrNode
func NewOrNode(left Node, right Node) OrNode {
	return OrNode{
		left:  left,
		right: right,
	}
}

// Childs returns collection of node childs
func (n OrNode) Childs() []Node {
	return []Node{n.left, n.right}
}

// Accept accepts translate visitor to invoke TranslateOrNode method
func (n OrNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateOrNode(n)
}
