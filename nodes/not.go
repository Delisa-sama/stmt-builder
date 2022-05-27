package nodes

type NotNode struct {
	orig Node
}

func NewNotNode(orig Node) NotNode {
	return NotNode{orig: orig}
}

func (n NotNode) Childs() []Node {
	return []Node{n.orig}
}

func (n NotNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNotNode(n)
}
