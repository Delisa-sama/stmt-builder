package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// NeOperator represents not equals operator
type NeOperator struct{}

// Node returns NeNode
func (o NeOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewNeNode(nodes.NewNameNode(leftOp), rightOp)
}
