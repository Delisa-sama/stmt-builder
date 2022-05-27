package nodes

type GtNode struct {
	name  NameNode
	right Node
}

func NewGtNode(name NameNode, right Node) GtNode {
	return GtNode{
		name:  name,
		right: right,
	}
}

func (n GtNode) Name() NameNode {
	return n.name
}

func (n GtNode) Childs() []Node {
	return []Node{n.right}
}

func (n GtNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateGtNode(n)
}
