package nodes

// ILikeNode represents equals operator
type ILikeNode struct {
	name  NameNode
	right Node
}

// NewILikeNode returns new ILikeNode
func NewILikeNode(name NameNode, right Node) ILikeNode {
	return ILikeNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n ILikeNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n ILikeNode) Childs() []Node {
	return []Node{n.right}
}

// Right returns child
func (n ILikeNode) Right() Node {
	return n.right
}

// Accept accepts translate visitor to invoke TranslateILikeNode method
func (n ILikeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateILikeNode(n)
}
