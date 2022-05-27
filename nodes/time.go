package nodes

import (
	"time"
)

type TimeNode struct {
	time.Time
}

func NewTimeNode(value time.Time) TimeNode {
	return TimeNode{value}
}

func (n TimeNode) Accept(visitor TranslateVisitor) string {
	return visitor.TranslateTimeNode(n)
}
