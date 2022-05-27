package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type LTOperator struct{}

func (o LTOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLtNode(nodes.NewNameNode(leftOp), rightOp)
}
