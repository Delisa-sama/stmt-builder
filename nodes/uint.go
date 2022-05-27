package nodes

// UintNode represents node with unsigned integer value
type UintNode uint64

// NewUintNode returns new UintNode
func NewUintNode(value uint64) UintNode {
	return UintNode(value)
}

// Accept accepts translate visitor to invoke TranslateUintNode method
func (n UintNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateUintNode(n)
}
