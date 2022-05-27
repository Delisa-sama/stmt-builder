package nodes

// AndNode represents AND operator
type AndNode struct {
	left  Node
	right Node
}

// NewAndNode returns new AndNode
func NewAndNode(left Node, right Node) AndNode {
	return AndNode{
		left:  left,
		right: right,
	}
}

// Childs returns collection of childs
func (n AndNode) Childs() []Node {
	return []Node{n.left, n.right}
}

// Accept accepts translate visitor to invoke TranslateAndNode method
func (n AndNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateAndNode(n)
}
