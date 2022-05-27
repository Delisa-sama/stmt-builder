package nodes

type StringNode string

func NewStringNode(value string) StringNode {
	return StringNode(value)
}

func (n StringNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateStringNode(n)
}
