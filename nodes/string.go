package nodes

// StringNode represents node with string value
type StringNode string

// NewStringNode returns new StringNode
func NewStringNode(value string) StringNode {
	return StringNode(value)
}

// Accept accepts translate visitor to invoke TranslateStringNode method
func (n StringNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateStringNode(n)
}
