package nodes

// NotNode represents negotiate node
type NotNode struct {
	orig Node
}

// NewNotNode returns new NotNode
func NewNotNode(orig Node) NotNode {
	return NotNode{orig: orig}
}

// Childs returns collection of childs
func (n NotNode) Childs() []Node {
	return []Node{n.orig}
}

// Accept accepts translate visitor to invoke TranslateNotNode method
func (n NotNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNotNode(n)
}
