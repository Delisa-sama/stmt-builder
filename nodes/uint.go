package nodes

type UintNode uint64

func NewUintNode(value uint64) UintNode {
	return UintNode(value)
}

func (n UintNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateUintNode(n)
}
