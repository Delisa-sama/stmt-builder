package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type lt struct {
	value SingleValue
}

// LT represents larger than operator
func LT(value SingleValue) statement.Operator {
	return lt{value: value}
}

// Node returns LtNode
func (o lt) Node(leftOp string) nodes.Node {
	return nodes.NewLtNode(nodes.NewNameNode(leftOp), o.value.Node())
}
