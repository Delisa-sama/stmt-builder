package nodes

type Node interface {
	Accept(visitor TranslateVisitor) string
}

type NodeWithRight interface {
	Node
	Right() Node
}

type NodeWithChilds interface {
	Node
	Childs() []Node
}

type NodeWithName interface {
	Node
	Name() NameNode
}
