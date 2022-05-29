package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type in struct {
	value MultipleValues
}

// In represents set operator
func In(value MultipleValues) statement.Operator {
	return in{value: value}
}

// Node returns InNode
func (o in) Node(leftOp string) nodes.Node {
	return nodes.NewInNode(nodes.NewNameNode(leftOp), nodes.NewArrayNode(o.value.Node()))
}
