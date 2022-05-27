package operators

import (
	"github.com/Delisa-sama/stmt-builder/nodes"
)

// InOperator represents set operator
type InOperator struct{}

// Node returns InNode
func (o InOperator) Node(leftOp string, rightOp nodes.Node) nodes.Node {
	return nodes.NewInNode(nodes.NewNameNode(leftOp), rightOp)
}
