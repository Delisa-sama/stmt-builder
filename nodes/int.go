package nodes

type IntNode int64

func NewIntNode(value int64) IntNode {
	return IntNode(value)
}

func (n IntNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateIntNode(n)
}
