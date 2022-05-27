package nodes

import (
	"time"
)

type TimeNode struct {
	value time.Time
}

func NewTimeNode(value time.Time) TimeNode {
	return TimeNode{value: value}
}

func (n TimeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateTimeNode(n)
}

func (n TimeNode) Time() time.Time {
	return n.value
}
