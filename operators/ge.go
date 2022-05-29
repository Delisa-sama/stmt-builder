package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type ge struct {
	value SingleValue
}

// GE represents greater or equals operator
func GE(value SingleValue) statement.Operator {
	return ge{value: value}
}

// Node returns GENode
func (o ge) Node(leftOp string) nodes.Node {
	return nodes.NewGENode(nodes.NewNameNode(leftOp), o.value.Node())
}
