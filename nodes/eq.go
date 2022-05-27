package nodes

type EqNode struct {
	name  NameNode
	right Node
}

func NewEqNode(name NameNode, right Node) EqNode {
	return EqNode{
		name:  name,
		right: right,
	}
}

func (n EqNode) Name() NameNode {
	return n.name
}

func (n EqNode) Childs() []Node {
	return []Node{n.right}
}

func (n EqNode) Right() Node {
	return n.right
}

func (n EqNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateEqNode(n)
}
