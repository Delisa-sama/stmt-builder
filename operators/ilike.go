package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
	"github.com/Delisa-sama/stmt-builder/statement"
)

type ilike struct {
	value string
}

// ILIKE represents ILIKE operator
func ILIKE(value string) statement.Operator {
	return ilike{value: value}
}

// Node returns EqNode
func (o ilike) Node(leftOp string) nodes.Node {
	return nodes.NewILikeNode(nodes.NewNameNode(leftOp), nodes.NewStringNode(o.value))
}
