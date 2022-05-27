package nodes

type LeNode struct {
	name  NameNode
	right Node
}

func NewLeNode(name NameNode, right Node) LeNode {
	return LeNode{
		name:  name,
		right: right,
	}
}

func (n LeNode) Name() NameNode {
	return n.name
}

func (n LeNode) Childs() []Node {
	return []Node{n.right}
}

func (n LeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateLeNode(n)
}
