package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type eq struct {
	value SingleValue
}

// EQ represents equals operator
func EQ(value SingleValue) statement.Operator {
	return eq{value: value}
}

// Node returns EqNode
func (o eq) Node(leftOp string) nodes.Node {
	return nodes.NewEqNode(nodes.NewNameNode(leftOp), o.value.Node())
}
