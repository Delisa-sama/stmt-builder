package nodes

// LikeNode represents equals operator
type LikeNode struct {
	name  NameNode
	right Node
}

// NewLikeNode returns new LikeNode
func NewLikeNode(name NameNode, right Node) LikeNode {
	return LikeNode{
		name:  name,
		right: right,
	}
}

// Name returns name
func (n LikeNode) Name() NameNode {
	return n.name
}

// Childs returns collection of childs
func (n LikeNode) Childs() []Node {
	return []Node{n.right}
}

// Right returns child
func (n LikeNode) Right() Node {
	return n.right
}

// Accept accepts translate visitor to invoke TranslateLikeNode method
func (n LikeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateLikeNode(n)
}
