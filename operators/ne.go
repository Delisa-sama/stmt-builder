package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type NeOperator struct{}

func (o NeOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NeNode{
		Left:  nodes.NameNode{Name: leftOp},
		Right: rightOp,
	}
}
