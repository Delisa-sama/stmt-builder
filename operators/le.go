package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type le struct {
	value SingleValue
}

// LE represents larger or equals operator
func LE(value SingleValue) statement.Operator {
	return le{value: value}
}

// Node returns LeNode
func (o le) Node(leftOp string) nodes.Node {
	return nodes.NewLeNode(nodes.NewNameNode(leftOp), o.value.Node())
}
