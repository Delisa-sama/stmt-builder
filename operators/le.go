package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type LEOperator struct{}

func (o LEOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLeNode(nodes.NewNameNode(leftOp), rightOp)
}
