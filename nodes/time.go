package nodes

import (
	"time"
)

// TimeNode represents node with time.Time value
type TimeNode struct {
	time.Time
}

// NewTimeNode returns new TimeNode
func NewTimeNode(value time.Time) TimeNode {
	return TimeNode{value}
}

// Accept accepts translate visitor to invoke TranslateTimeNode method
func (n TimeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateTimeNode(n)
}

// Value returns value of the node as primitive type
func (n TimeNode) Value() interface{} {
	return n.Time
}
