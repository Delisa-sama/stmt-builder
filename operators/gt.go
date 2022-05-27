package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type GTOperator struct{}

func (o GTOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGtNode(nodes.NewNameNode(leftOp), rightOp)
}
