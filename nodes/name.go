package nodes

// NameNode represents some name
type NameNode struct {
	name string
}

// NewNameNode returns new NameNode
func NewNameNode(name string) NameNode {
	return NameNode{name: name}
}

// Accept accepts translate visitor to invoke TranslateNameNode method
func (n NameNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNameNode(n)
}

// Name returns name
func (n NameNode) Name() string {
	return n.name
}
