package nodes

type InNode struct {
	name  NameNode
	right Node
}

func NewInNode(name NameNode, right Node) InNode {
	return InNode{
		name:  name,
		right: right,
	}
}

func (n InNode) Name() NameNode {
	return n.name
}

func (n InNode) Childs() []Node {
	return []Node{n.right}
}

func (n InNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateInNode(n)
}
