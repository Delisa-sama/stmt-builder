package nodes

type GeNode struct {
	name  NameNode
	right Node
}

func NewGENode(name NameNode, right Node) GeNode {
	return GeNode{
		name:  name,
		right: right,
	}
}

func (n GeNode) Name() NameNode {
	return n.name
}

func (n GeNode) Childs() []Node {
	return []Node{n.right}
}

func (n GeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateGeNode(n)
}
