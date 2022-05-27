package nodes

type AndNode struct {
	left  Node
	right Node
}

func NewAndNode(left Node, right Node) AndNode {
	return AndNode{
		left:  left,
		right: right,
	}
}

func (n AndNode) Childs() []Node {
	return []Node{n.left, n.right}
}

func (n AndNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateAndNode(n)
}
