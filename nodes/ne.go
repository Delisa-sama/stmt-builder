package nodes

// NeNode represents not equals node
type NeNode struct {
	name  NameNode
	right Node
}

// NewNeNode returns new NeNode
func NewNeNode(name NameNode, right Node) NeNode {
	return NeNode{
		name:  name,
		right: right,
	}
}

// Name return nodes name
func (n NeNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n NeNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateNeNode method
func (n NeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNeNode(n)
}

// Right returns the right child
func (n NeNode) Right() Node {
	return n.right
}
