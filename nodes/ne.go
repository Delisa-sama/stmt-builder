package nodes

type NeNode struct {
	name  NameNode
	right Node
}

func NewNeNode(name NameNode, right Node) NeNode {
	return NeNode{
		name:  name,
		right: right,
	}
}

func (n NeNode) Name() NameNode {
	return n.name
}

func (n NeNode) Childs() []Node {
	return []Node{n.right}
}

func (n NeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNeNode(n)
}

func (n NeNode) Right() Node {
	return n.right
}
