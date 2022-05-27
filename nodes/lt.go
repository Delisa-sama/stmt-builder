package nodes

// LtNode represents less than operator
type LtNode struct {
	name  NameNode
	right Node
}

// NewLtNode returns new LtNode
func NewLtNode(name NameNode, right Node) LtNode {
	return LtNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n LtNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n LtNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateLtNode method
func (n LtNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateLtNode(n)
}
