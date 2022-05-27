package nodes

type NameNode struct {
	name string
}

func NewNameNode(name string) NameNode {
	return NameNode{name: name}
}

func (n NameNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNameNode(n)
}

func (n NameNode) Name() string {
	return n.name
}
