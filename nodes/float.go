package nodes

// FloatNode represents node with float value
type FloatNode float64

// NewFloatNode returns new FloatNode
func NewFloatNode(value float64) FloatNode {
	return FloatNode(value)
}

// Accept accepts translate visitor to invoke TranslateFloatNode method
func (n FloatNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateFloatNode(n)
}
