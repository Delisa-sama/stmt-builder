package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type like struct {
	value string
}

// LIKE represents LIKE operator
func LIKE(value string) statement.Operator {
	return like{value: value}
}

// Node returns EqNode
func (o like) Node(leftOp string) nodes.Node {
	return nodes.NewLikeNode(nodes.NewNameNode(leftOp), nodes.NewStringNode(o.value))
}
