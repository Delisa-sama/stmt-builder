package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type NeOperator struct{}

func (o NeOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewNeNode(nodes.NewNameNode(leftOp), rightOp)
}
