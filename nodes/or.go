package nodes

type OrNode struct {
	left  Node
	right Node
}

func NewOrNode(left Node, right Node) OrNode {
	return OrNode{
		left:  left,
		right: right,
	}
}

func (n OrNode) Childs() []Node {
	return []Node{n.left, n.right}
}

func (n OrNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateOrNode(n)
}
