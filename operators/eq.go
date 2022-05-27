package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type EQOperator struct{}

func (o EQOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewEqNode(nodes.NewNameNode(leftOp), rightOp)
}
