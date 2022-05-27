package nodes

type FloatNode float64

func NewFloatNode(value float64) FloatNode {
	return FloatNode(value)
}

func (n FloatNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateFloatNode(n)
}
