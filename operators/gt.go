package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type gt struct {
	value SingleValue
}

// GT represents greater than operator
func GT(value SingleValue) statement.Operator {
	return gt{value: value}
}

// Node returns GtNode
func (o gt) Node(leftOp string) nodes.Node {
	return nodes.NewGtNode(nodes.NewNameNode(leftOp), o.value.Node())
}
