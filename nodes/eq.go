package nodes

// EqNode represents equals operator
type EqNode struct {
	name  NameNode
	right Node
}

// NewEqNode returns new EqNode
func NewEqNode(name NameNode, right Node) EqNode {
	return EqNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n EqNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n EqNode) Childs() []Node {
	return []Node{n.right}
}

// Right returns child
func (n EqNode) Right() Node {
	return n.right
}

// Accept accepts translate visitor to invoke TranslateEqNode method
func (n EqNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateEqNode(n)
}
