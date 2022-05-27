package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

type GEOperator struct{}

func (o GEOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewGENode(nodes.NewNameNode(leftOp), rightOp)
}
