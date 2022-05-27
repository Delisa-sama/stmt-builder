package nodes

// GtNode represents greater than operator
type GtNode struct {
	name  NameNode
	right Node
}

// NewGtNode returns new GtNode
func NewGtNode(name NameNode, right Node) GtNode {
	return GtNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n GtNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n GtNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateGtNode method
func (n GtNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateGtNode(n)
}
