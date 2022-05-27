package nodes

type NullNode struct{}

func NewNullNode() NullNode {
	return NullNode{}
}

func (n NullNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateNullNode(n)
}
