package nodes

type LtNode struct {
	name  NameNode
	right Node
}

func NewLtNode(name NameNode, right Node) LtNode {
	return LtNode{
		name:  name,
		right: right,
	}
}

func (n LtNode) Name() NameNode {
	return n.name
}

func (n LtNode) Childs() []Node {
	return []Node{n.right}
}

func (n LtNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateLtNode(n)
}
