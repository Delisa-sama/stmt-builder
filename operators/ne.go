package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type ne struct {
	value SingleValue
}

// NE represents not equals operator
func NE(value SingleValue) statement.Operator {
	return ne{value: value}
}

// Node returns NeNode
func (o ne) Node(leftOp string) nodes.Node {
	return nodes.NewNeNode(nodes.NewNameNode(leftOp), o.value.Node())
}
