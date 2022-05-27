package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// LTOperator represents larger than operator
type LTOperator struct{}

// Node returns LtNode
func (o LTOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewLtNode(nodes.NewNameNode(leftOp), rightOp)
}
