package nodes

// GeNode represents greater or equals operator
type GeNode struct {
	name  NameNode
	right Node
}

// NewGENode returns GeNode
func NewGENode(name NameNode, right Node) GeNode {
	return GeNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n GeNode) Name() NameNode {
	return n.name
}

// Childs returns childs
func (n GeNode) Childs() []Node {
	return []Node{n.right}
}

// Accept accepts translate visitor to invoke TranslateGeNode method
func (n GeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateGeNode(n)
}
