package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type notIn struct {
	value MultipleValues
}

// NotIn represents set operator
func NotIn(value MultipleValues) statement.Operator {
	return notIn{value: value}
}

// Node returns InNode
func (o notIn) Node(leftOp string) nodes.Node {
	return nodes.NewNotInNode(nodes.NewNameNode(leftOp), nodes.NewArrayNode(o.value.Node()))
}
