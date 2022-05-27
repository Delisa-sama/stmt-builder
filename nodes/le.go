package nodes

// LeNode represents less of equals operator
type LeNode struct {
	name  NameNode
	right Node
}

// NewLeNode returns LeNode
func NewLeNode(name NameNode, right Node) LeNode {
	return LeNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n LeNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n LeNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateLeNode method
func (n LeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateLeNode(n)
}
