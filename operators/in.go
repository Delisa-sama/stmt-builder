package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type InOperator struct{}

func (o InOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewInNode(nodes.NewNameNode(leftOp), rightOp)
}
