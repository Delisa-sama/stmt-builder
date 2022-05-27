package nodes

// InNode represents set operator
type InNode struct {
	name  NameNode
	right Node
}

// NewInNode returns InNode
func NewInNode(name NameNode, right Node) InNode {
	return InNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n InNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n InNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateInNode method
func (n InNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateInNode(n)
}
